package app

import (
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/event"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/params"
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

	Blockdb  db.DBAccessor
	Objectdb mapobjectdb.DBAccessor

	BlockAccessor    *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
	Mapblockrenderer *mapblockrenderer.MapBlockRenderer
	Tilerenderer     *tilerenderer.TileRenderer
	Events           event.EventConsumer
}
