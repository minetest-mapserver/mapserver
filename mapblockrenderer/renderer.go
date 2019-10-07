package mapblockrenderer

import (
	"errors"
	"image"
	"image/color"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/mapblockaccessor"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type MapBlockRenderer struct {
	accessor           *mapblockaccessor.MapBlockAccessor
	colors             *colormapping.ColorMapping
	enableShadow       bool
	enableTransparency bool
}

func NewMapBlockRenderer(accessor *mapblockaccessor.MapBlockAccessor, colors *colormapping.ColorMapping) *MapBlockRenderer {
	return &MapBlockRenderer{
		accessor:           accessor,
		colors:             colors,
		enableShadow:       true,
		enableTransparency: false,
	}
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

func addColorComponent(c *color.RGBA, value int) *color.RGBA {
	return &color.RGBA{
		R: clamp(int(c.R) + value),
		G: clamp(int(c.G) + value),
		B: clamp(int(c.B) + value),
		A: c.A,
	}
}

func (r *MapBlockRenderer) Render(pos1, pos2 *coords.MapBlockCoords) (*image.NRGBA, error) {
	if pos1.X != pos2.X {
		return nil, errors.New("X does not line up")
	}

	if pos1.Z != pos2.Z {
		return nil, errors.New("Z does not line up")
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
		currentPos := coords.NewMapBlockCoords(pos1.X, mapBlockY, pos1.Z)
		mb, err := r.accessor.GetMapBlock(currentPos)

		if err != nil {
			return nil, err
		}

		if mb == nil || mb.IsEmpty() {
			continue
		}

		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				for y := 15; y >= 0; y-- {
					if xzOccupationMap[x][z] {
						break
					}

					nodeName := mb.GetNodeName(x, y, z)
					param2 := mb.GetParam2(x, y, z)

					if nodeName == "" {
						continue
					}

					c := r.colors.GetColor(nodeName, param2)

					if c == nil {
						continue
					}

					if r.enableShadow {
						var left, leftAbove, top, topAbove string

						if x > 0 {
							//same mapblock
							left = mb.GetNodeName(x-1, y, z)
							if y < 15 {
								leftAbove = mb.GetNodeName(x-1, y+1, z)
							}

						} else {
							//neighbouring mapblock
							neighbourPos := coords.NewMapBlockCoords(currentPos.X-1, currentPos.Y, currentPos.Z)
							neighbourMapblock, err := r.accessor.GetMapBlock(neighbourPos)

							if neighbourMapblock != nil && err == nil {
								left = neighbourMapblock.GetNodeName(15, y, z)
								if y < 15 {
									leftAbove = neighbourMapblock.GetNodeName(15, y+1, z)
								}
							}
						}

						if z < 14 {
							//same mapblock
							top = mb.GetNodeName(x, y, z+1)
							if y < 15 {
								topAbove = mb.GetNodeName(x, y+1, z+1)
							}

						} else {
							//neighbouring mapblock
							neighbourPos := coords.NewMapBlockCoords(currentPos.X, currentPos.Y, currentPos.Z+1)
							neighbourMapblock, err := r.accessor.GetMapBlock(neighbourPos)

							if neighbourMapblock != nil && err == nil {
								top = neighbourMapblock.GetNodeName(x, y, 0)
								if y < 15 {
									topAbove = neighbourMapblock.GetNodeName(x, y+1, 0)
								}
							}
						}

						if IsViewBlocking(leftAbove) {
							//add shadow
							c = addColorComponent(c, -10)
						}

						if IsViewBlocking(topAbove) {
							//add shadow
							c = addColorComponent(c, -10)
						}

						if !IsViewBlocking(left) {
							//add light
							c = addColorComponent(c, 10)
						}

						if !IsViewBlocking(top) {
							//add light
							c = addColorComponent(c, 10)
						}
					}

					imgX := x * IMG_SCALE
					imgY := (15 - z) * IMG_SCALE

					r32, g32, b32, a32 := c.RGBA()
					r8, g8, b8, a8 := uint8(r32), uint8(g32), uint8(b32), uint8(a32)
					for Y := imgY; Y < imgY+IMG_SCALE; Y++ {
						ix := (Y*IMG_SIZE + imgX) << 2
						for X := 0; X < IMG_SCALE; X++ {
							img.Pix[ix] = r8
							ix++
							img.Pix[ix] = g8
							ix++
							img.Pix[ix] = b8
							ix++
							img.Pix[ix] = a8
							ix++
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
