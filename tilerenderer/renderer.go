package tilerenderer

import (
  "mapserver/coords"
  "mapserver/mapblockrenderer"
  "mapserver/tiledb"
  "image"
)

type TileRenderer struct {
  mapblockrenderer *mapblockrenderer.MapBlockRenderer
  tiledb *tiledb.DBAccessor
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer, tiledb *tiledb.DBAccessor) *TileRenderer {
  return &TileRenderer{
    mapblockrenderer: mapblockrenderer,
    tiledb: tiledb,
  }
}

//TODO layerConfig
func (tr *TileRenderer) Render(tc coords.TileCoords, layerId int) (*image.NRGBA, error) {
  return nil, nil
}
