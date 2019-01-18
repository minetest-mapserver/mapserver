package app

import (
  "mapserver/params"
  "mapserver/worldconfig"
  "mapserver/db"
  "mapserver/mapblockaccessor"
  "mapserver/colormapping"
  "mapserver/mapblockrenderer"
  "mapserver/tiledb"
  "mapserver/tilerenderer"
)

const (
	Version = "2.0-DEV"
)

type App struct {
  Params params.ParamsType
  Config *Config
  Worldconfig worldconfig.WorldConfig

  Blockdb db.DBAccessor
  Tiledb tiledb.DBAccessor

  BlockAccessor *mapblockaccessor.MapBlockAccessor
  Colormapping *colormapping.ColorMapping
  Mapblockrenderer *mapblockrenderer.MapBlockRenderer
  Tilerenderer *tilerenderer.TileRenderer
}
