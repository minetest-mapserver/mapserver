package mapblockrenderer

import (
	"errors"
	"image"
	"image/color"
	"mapserver/mapblockaccessor"
	"mapserver/types"
	"strings"
	"sync"
	"time"

	"github.com/minetest-go/colormapping"
	"github.com/minetest-go/mapparser"
	mtypes "github.com/minetest-go/types"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type MapBlockRenderer struct {
	accessor           *mapblockaccessor.MapBlockAccessor
	colors             *colormapping.ColorMapping
	enableShadow       bool
	enableTransparency bool

	// node names colored by param2 (colormapping's extended palette)
	paletteNodes map[string]bool
	// resolved colors by node name, including misses (the colormapping does not cache those)
	colorCache map[string]nodeColor
	colorMutex sync.RWMutex
}

// color of a node name
type nodeColor struct {
	color    color.RGBA
	hasColor bool
	// color depends on param2, resolved per node
	palette bool
}

// per-mapblock node info, indexed by content id
type nodeInfo struct {
	nodeColor
	viewBlocking bool
}

func NewMapBlockRenderer(accessor *mapblockaccessor.MapBlockAccessor, colors *colormapping.ColorMapping) *MapBlockRenderer {
	return &MapBlockRenderer{
		accessor:           accessor,
		colors:             colors,
		enableShadow:       true,
		enableTransparency: false,
		paletteNodes:       loadPaletteNodes(),
		colorCache:         map[string]nodeColor{},
	}
}

// loads the node names the colormapping colors by param2, the list is not exported otherwise
func loadPaletteNodes() map[string]bool {
	nodes := map[string]bool{}
	data, err := colormapping.Files.ReadFile("colors/extended_palette.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		name := strings.TrimSpace(line)
		if name != "" {
			nodes[name] = true
		}
	}
	return nodes
}

// returns the color of the node name, cached
func (r *MapBlockRenderer) getNodeColor(name string) nodeColor {
	r.colorMutex.RLock()
	nc, found := r.colorCache[name]
	r.colorMutex.RUnlock()
	if found {
		return nc
	}

	if r.paletteNodes[name] {
		nc = nodeColor{palette: true}
	} else if c := r.colors.GetColor(name, 0); c != nil {
		// copy, the returned color is shared
		nc = nodeColor{color: *c, hasColor: true}
	}

	r.colorMutex.Lock()
	r.colorCache[name] = nc
	r.colorMutex.Unlock()
	return nc
}

// builds the content-id indexed node info table of a mapblock
func (r *MapBlockRenderer) getNodeInfos(mb *mtypes.MapBlock) []nodeInfo {
	maxId := -1
	for id := range mb.BlockMapping {
		if id > maxId {
			maxId = id
		}
	}

	infos := make([]nodeInfo, maxId+1)
	for id, name := range mb.BlockMapping {
		if id < 0 || name == "" {
			continue
		}
		infos[id] = nodeInfo{
			nodeColor:    r.getNodeColor(name),
			viewBlocking: IsViewBlocking(name),
		}
	}
	return infos
}

// a mapblock with its node info table
type renderBlock struct {
	mb    *mtypes.MapBlock
	infos []nodeInfo
}

var emptyNodeInfo = nodeInfo{}

// returns the node info at the mapblock-relative position, the empty info for unknown or missing nodes
func (b *renderBlock) nodeAt(x, y, z int) *nodeInfo {
	if b == nil {
		return &emptyNodeInfo
	}
	pos := mapparser.GetNodePos(x, y, z)
	if pos >= len(b.mb.ContentId) {
		return &emptyNodeInfo
	}
	id := b.mb.ContentId[pos]
	if id < 0 || id >= len(b.infos) {
		return &emptyNodeInfo
	}
	return &b.infos[id]
}

// lazily loaded neighbour mapblock, used for the shading at the mapblock edges
type neighbourBlock struct {
	loaded bool
	block  *renderBlock
}

func (r *MapBlockRenderer) getNeighbour(n *neighbourBlock, pos *types.MapBlockCoords) *renderBlock {
	if !n.loaded {
		n.loaded = true
		mb, err := r.accessor.GetMapBlock(pos)
		if mb != nil && err == nil {
			n.block = &renderBlock{mb: mb, infos: r.getNodeInfos(mb)}
		}
	}
	return n.block
}

const (
	IMG_SCALE                         = 16
	IMG_SIZE                          = IMG_SCALE * 16
	EXPECTED_BLOCKS_PER_FLAT_MAPBLOCK = 16 * 16
)

func IsViewBlocking(nodeName string) bool {
	if nodeName == "" {
		return false
	}

	if nodeName == "vacuum:vacuum" {
		return false
	}

	if nodeName == "air" {
		return false
	}

	return true
}

func clamp(num int) uint8 {
	if num < 0 {
		return 0
	}

	if num > 255 {
		return 255
	}

	return uint8(num)
}

func addColorComponent(c color.RGBA, value int) color.RGBA {
	return color.RGBA{
		R: clamp(int(c.R) + value),
		G: clamp(int(c.G) + value),
		B: clamp(int(c.B) + value),
		A: c.A,
	}
}

func (r *MapBlockRenderer) Render(pos1, pos2 *types.MapBlockCoords) (*image.NRGBA, error) {
	if pos1.X != pos2.X {
		return nil, errors.New("x does not line up")
	}

	if pos1.Z != pos2.Z {
		return nil, errors.New("z does not line up")
	}

	renderedMapblocks.Inc()
	timer := prometheus.NewTimer(renderDuration)
	defer timer.ObserveDuration()

	start := time.Now()
	defer func() {
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Debug("Rendering completed")
	}()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMG_SIZE, IMG_SIZE}
	img := image.NewNRGBA(image.Rectangle{upLeft, lowRight})

	maxY := pos1.Y
	minY := pos2.Y

	if minY > maxY {
		maxY, minY = minY, maxY
	}

	foundBlocks := 0
	xzOccupationMap := make([][]bool, 16)
	for x := range xzOccupationMap {
		xzOccupationMap[x] = make([]bool, 16)
	}

	for mapBlockY := maxY; mapBlockY >= minY; mapBlockY-- {
		currentPos := types.NewMapBlockCoords(pos1.X, mapBlockY, pos1.Z)
		mb, err := r.accessor.GetMapBlock(currentPos)

		if err != nil {
			return nil, err
		}

		if mb == nil || mb.IsEmpty() {
			continue
		}

		block := &renderBlock{mb: mb, infos: r.getNodeInfos(mb)}
		left := &neighbourBlock{}
		top := &neighbourBlock{}

		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				for y := 15; y >= 0; y-- {
					if xzOccupationMap[x][z] {
						break
					}

					node := block.nodeAt(x, y, z)

					var c color.RGBA
					if node.palette {
						nodePos := mapparser.GetNodePos(x, y, z)
						pc := r.colors.GetColor(mb.BlockMapping[mb.ContentId[nodePos]], mb.Param2[nodePos])
						if pc == nil {
							continue
						}
						c = *pc
					} else if node.hasColor {
						c = node.color
					} else {
						continue
					}

					// clamp alpha channel to max
					c.A = 255

					if r.enableShadow {
						var leftNode, leftAbove, topNode, topAbove *nodeInfo = &emptyNodeInfo, &emptyNodeInfo, &emptyNodeInfo, &emptyNodeInfo

						if x > 0 {
							//same mapblock
							leftNode = block.nodeAt(x-1, y, z)
							if y < 15 {
								leftAbove = block.nodeAt(x-1, y+1, z)
							}

						} else {
							//neighbouring mapblock
							nb := r.getNeighbour(left, types.NewMapBlockCoords(currentPos.X-1, currentPos.Y, currentPos.Z))
							leftNode = nb.nodeAt(15, y, z)
							if y < 15 {
								leftAbove = nb.nodeAt(15, y+1, z)
							}
						}

						if z < 15 {
							//same mapblock
							topNode = block.nodeAt(x, y, z+1)
							if y < 15 {
								topAbove = block.nodeAt(x, y+1, z+1)
							}

						} else {
							//neighbouring mapblock
							nb := r.getNeighbour(top, types.NewMapBlockCoords(currentPos.X, currentPos.Y, currentPos.Z+1))
							topNode = nb.nodeAt(x, y, 0)
							if y < 15 {
								topAbove = nb.nodeAt(x, y+1, 0)
							}
						}

						if leftAbove.viewBlocking {
							//add shadow
							c = addColorComponent(c, -10)
						}

						if topAbove.viewBlocking {
							//add shadow
							c = addColorComponent(c, -10)
						}

						if !leftNode.viewBlocking {
							//add light
							c = addColorComponent(c, 10)
						}

						if !topNode.viewBlocking {
							//add light
							c = addColorComponent(c, 10)
						}
					}

					imgX := x * IMG_SCALE
					imgY := (15 - z) * IMG_SCALE

					// alpha is always 255: premultiplied and non-premultiplied values are the same
					for Y := imgY; Y < imgY+IMG_SCALE; Y++ {
						ix := (Y*IMG_SIZE + imgX) << 2
						for X := 0; X < IMG_SCALE; X++ {
							img.Pix[ix] = c.R
							img.Pix[ix+1] = c.G
							img.Pix[ix+2] = c.B
							img.Pix[ix+3] = c.A
							ix += 4
						}
					}

					if c.A != 0xFF || !r.enableTransparency {
						//not transparent, mark as rendered
						foundBlocks++
						xzOccupationMap[x][z] = true
					}

					if foundBlocks == EXPECTED_BLOCKS_PER_FLAT_MAPBLOCK {
						return img, nil
					}
				}
			}
		}
	}

	if foundBlocks == 0 {
		return nil, nil
	}

	return img, nil
}
