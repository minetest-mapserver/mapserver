package mapblockrenderer

import (
  "mapserver/coords"
  "mapserver/db"
  "mapserver/colormapping"
)

type MapBlockRenderer struct {
  accessor *coords.DBAccessor
  colors *colormapping.ColorMapping
}

func NewMapBlockRenderer(accessor *coords.DBAccessor, colors *colormapping.ColorMapping){
  //TODO
}

func (r *MapBlockRenderer) Render(range coords.MapBlockRange){
  //TODO
}
