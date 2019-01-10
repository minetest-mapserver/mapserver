package mapblockrenderer

import (
	"mapserver/colormapping"
	"mapserver/coords"
)

type MapBlockRenderer struct {
	accessor *coords.DBAccessor
	colors   *colormapping.ColorMapping
}

func NewMapBlockRenderer(accessor *coords.DBAccessor, colors *colormapping.ColorMapping) {
	//TODO
}

func (r *MapBlockRenderer) Render(r coords.MapBlockRange) {
	//TODO
}
