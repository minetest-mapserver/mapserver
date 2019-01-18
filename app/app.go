package app

import (
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/params"
	"mapserver/tiledb"
	"mapserver/tilerenderer"
	"mapserver/worldconfig"
)

const (
	Version = "2.0-DEV"
)

type App struct {
	Params      params.ParamsType
	Config      *Config
	Worldconfig worldconfig.WorldConfig

	Blockdb db.DBAccessor
	Tiledb  tiledb.DBAccessor

	BlockAccessor    *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
	Mapblockrenderer *mapblockrenderer.MapBlockRenderer
	Tilerenderer     *tilerenderer.TileRenderer
}
