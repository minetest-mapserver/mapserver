package mapblockrenderer

import (
	"errors"
	"github.com/sirupsen/logrus"
	"image"
	"image/draw"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/mapblockaccessor"
	"time"
)

type MapBlockRenderer struct {
	accessor *mapblockaccessor.MapBlockAccessor
	colors   *colormapping.ColorMapping
}

func NewMapBlockRenderer(accessor *mapblockaccessor.MapBlockAccessor, colors *colormapping.ColorMapping) MapBlockRenderer {
	return MapBlockRenderer{accessor: accessor, colors: colors}
}

const (
	IMG_SCALE                         = 16
	IMG_SIZE                          = IMG_SCALE * 16
	EXPECTED_BLOCKS_PER_FLAT_MAPBLOCK = 16 * 16
)

func (r *MapBlockRenderer) Render(pos1, pos2 coords.MapBlockCoords) (*image.NRGBA, error) {
	if pos1.X != pos2.X {
		return nil, errors.New("X does not line up")
	}

	if pos1.Z != pos2.Z {
		return nil, errors.New("Z does not line up")
	}

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

		if mb == nil {
			continue
		}

		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				for y := 15; y >= 0; y-- {
					if xzOccupationMap[x][z] {
						continue
					}

					nodeName := mb.GetNodeName(x, y, z)

					if nodeName == "" {
						continue
					}

					c := r.colors.GetColor(nodeName)

					if c == nil {
						continue
					}

					rect := image.Rect(
						x*IMG_SCALE, z*IMG_SCALE,
						(x*IMG_SCALE)+IMG_SCALE, (z*IMG_SCALE)+IMG_SCALE,
					)

					foundBlocks++
					xzOccupationMap[x][z] = true
					draw.Draw(img, rect, &image.Uniform{c}, image.ZP, draw.Src)

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
