package mapblockrenderer

import (
	"image/color"
	"mapserver/mapblockaccessor"
	"mapserver/types"
	"testing"
	"time"

	"github.com/minetest-go/colormapping"
	"github.com/minetest-go/mapparser"
	mtypes "github.com/minetest-go/types"
)

const (
	testAirId   = 0
	testStoneId = 1
)

// creates an air-filled mapblock
func newTestMapblock() *mtypes.MapBlock {
	mb := mtypes.NewMapblock()
	mb.ContentId = make([]int, 4096)
	mb.Param1 = make([]int, 4096)
	mb.Param2 = make([]int, 4096)
	mb.BlockMapping[testAirId] = "air"
	mb.BlockMapping[testStoneId] = "default:stone"
	return mb
}

func setTestNode(mb *mtypes.MapBlock, x, y, z, id int) {
	mb.ContentId[mapparser.GetNodePos(x, y, z)] = id
}

// the shading of a node depends on its neighbour in +z direction,
// which is in the same mapblock for z=0..14 and only for z=15 in the next one
func TestRenderShadowMapblockEdge(t *testing.T) {
	base := color.NRGBA{R: 100, G: 100, B: 100, A: 255}

	c := colormapping.NewColorMapping()
	_, err := c.LoadBytes([]byte("default:stone 100 100 100\n"))
	if err != nil {
		t.Fatal(err)
	}

	// flat stone floor with a single raised node at the +z edge
	mb := newTestMapblock()
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			setTestNode(mb, x, 0, z, testStoneId)
		}
	}
	setTestNode(mb, 5, 1, 15, testStoneId)

	// all blocks are served from the cache, the database is never queried
	accessor := mapblockaccessor.NewMapBlockAccessor(nil, time.Minute, 100)
	pos := types.NewMapBlockCoords(0, 0, 0)
	accessor.Update(pos, mb)
	// neighbours used for the shading at the -x and +z edges: not generated
	accessor.Update(types.NewMapBlockCoords(-1, 0, 0), nil)
	accessor.Update(types.NewMapBlockCoords(0, 0, 1), nil)

	r := NewMapBlockRenderer(accessor, c)
	img, err := r.Render(pos, pos)
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("no image rendered")
	}

	nodeColor := func(x, z int) color.NRGBA {
		return img.NRGBAAt(x*IMG_SCALE, (15-z)*IMG_SCALE)
	}
	shaded := func(value int) color.NRGBA {
		v := uint8(int(base.R) + value)
		return color.NRGBA{R: v, G: v, B: v, A: 255}
	}

	tests := []struct {
		name string
		x, z int
		want color.NRGBA
	}{
		{"inner floor", 10, 10, base},
		{"floor at z=14, neighbour in the same mapblock", 10, 14, base},
		{"floor at z=14, below the raised node", 5, 14, shaded(-10)},
		{"floor at z=15, neighbour mapblock not generated", 10, 15, shaded(10)},
		{"raised node at z=15", 5, 15, shaded(20)},
	}

	for _, tc := range tests {
		got := nodeColor(tc.x, tc.z)
		if got != tc.want {
			t.Errorf("%s: node %d/%d has color %v, expected %v", tc.name, tc.x, tc.z, got, tc.want)
		}
	}
}
