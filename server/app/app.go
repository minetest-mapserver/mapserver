package app

import (
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/params"
	"mapserver/settings"
	"mapserver/tilerenderer"
)

const (
	Version = "0.0.2"
)

type App struct {
	Params      params.ParamsType
	Config      *Config
	Worldconfig map[string]string

	Blockdb  db.DBAccessor
	Objectdb mapobjectdb.DBAccessor
	Settings *settings.Settings

	BlockAccessor    *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
	Mapblockrenderer *mapblockrenderer.MapBlockRenderer
	Tilerenderer     *tilerenderer.TileRenderer

	WebEventbus *eventbus.Eventbus
}
