package app

import (
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/params"
	"mapserver/tilerenderer"
	"mapserver/worldconfig"
)

const (
	Version = "0.0.1"
)

type App struct {
	Params      params.ParamsType
	Config      *Config
	Worldconfig worldconfig.WorldConfig

	Blockdb  db.DBAccessor
	Objectdb mapobjectdb.DBAccessor

	BlockAccessor    *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
	Mapblockrenderer *mapblockrenderer.MapBlockRenderer
	Tilerenderer     *tilerenderer.TileRenderer
}
