package app

import (
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/params"
	"mapserver/settings"
	"mapserver/tiledb"
	"mapserver/tilerenderer"

	"github.com/minetest-go/colormapping"
)

type App struct {
	Params      params.ParamsType
	Config      *Config
	Worldconfig map[string]string

	Blockdb  db.DBAccessor
	Objectdb mapobjectdb.DBAccessor
	TileDB   *tiledb.TileDB
	Settings settings.Settings

	MapBlockAccessor *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
	Mapblockrenderer *mapblockrenderer.MapBlockRenderer
	Tilerenderer     *tilerenderer.TileRenderer

	MediaRepo map[string][]byte

	WebEventbus *eventbus.Eventbus
}
