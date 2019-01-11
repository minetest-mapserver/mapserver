package mapblockrenderer

import (
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/mapblockaccessor"
)

type MapBlockRenderer struct {
	accessor *mapblockaccessor.MapBlockAccessor
	colors   *colormapping.ColorMapping
}

func NewMapBlockRenderer(accessor *mapblockaccessor.MapBlockAccessor, colors *colormapping.ColorMapping) MapBlockRenderer {
	return MapBlockRenderer{accessor: accessor, colors: colors}
}

func (r *MapBlockRenderer) Render(pos1, pos2 coords.MapBlockCoords) {
	//TODO
}
