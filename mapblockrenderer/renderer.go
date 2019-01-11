package mapblockrenderer

import (
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/mapblockaccessor"
	"image"
	"image/color"
)

type MapBlockRenderer struct {
	accessor *mapblockaccessor.MapBlockAccessor
	colors   *colormapping.ColorMapping
}

func NewMapBlockRenderer(accessor *mapblockaccessor.MapBlockAccessor, colors *colormapping.ColorMapping) MapBlockRenderer {
	return MapBlockRenderer{accessor: accessor, colors: colors}
}

const (
	IMG_SCALE = 16
	IMG_SIZE = IMG_SCALE * 16
)

func (r *MapBlockRenderer) Render(pos1, pos2 coords.MapBlockCoords) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMG_SIZE, IMG_SIZE}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}
	img.Set(10, 10, cyan)
	return img
}
