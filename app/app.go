package app

import (
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/settings"
	"mapserver/tiledb"
	"mapserver/tilerenderer"
	"mapserver/util"

	"github.com/minetest-go/colormapping"
)

type App struct {
	Params      util.ParamsType
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

	WebEventbus *util.Eventbus
}
