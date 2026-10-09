package app

import (
	"mapserver/db/postgres"
	"mapserver/db/sqlite"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	postgresobjdb "mapserver/mapobjectdb/postgres"
	sqliteobjdb "mapserver/mapobjectdb/sqlite"
	"mapserver/settings"
	"mapserver/tiledb"
	"mapserver/tilerenderer"
	"mapserver/util"
	"time"

	"github.com/minetest-go/colormapping"

	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"errors"
)

func Setup(p util.ParamsType, cfg *Config) *App {
	a := App{}
	a.Params = p
	a.Config = cfg
	a.WebEventbus = util.NewEventbus()

	//Parse world config
	a.Worldconfig = util.ParseWorldConfig(filepath.Join(a.Config.WorldPath, "world.mt"))
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mapserver")

	var err error

	switch a.Worldconfig[util.CONFIG_BACKEND] {
	case util.BACKEND_SQLITE3:
		map_path := filepath.Join(a.Config.WorldPath, "map.sqlite")

		// check if the database exists, otherwise abort (nothing to render/display)
		_, err := os.Stat(map_path)
		if os.IsNotExist(err) {
			panic("world-map does not exist, aborting")
		}

		// create a new sqlite-based blockdb instance
		a.Blockdb, err = sqlite.New(map_path)
		if err != nil {
			panic(err)
		}

	case util.BACKEND_POSTGRES:
		// create a new postgres based blockdb
		a.Blockdb, err = postgres.New(a.Worldconfig[util.CONFIG_PSQL_CONNECTION])
		if err != nil {
			panic(err)
		}

	default:
		panic(errors.New("map-backend not supported: " + a.Worldconfig[util.CONFIG_BACKEND]))
	}

	//migrate block db
	err = a.Blockdb.Migrate()
	if err != nil {
		panic(err)
	}

	//mapblock accessor
	expireDuration, err := time.ParseDuration(cfg.MapBlockAccessorCfg.Expiretime)
	if err != nil {
		panic(err)
	}

	purgeDuration, err := time.ParseDuration(cfg.MapBlockAccessorCfg.Purgetime)
	if err != nil {
		panic(err)
	}

	// mapblock accessor
	a.MapBlockAccessor = mapblockaccessor.NewMapBlockAccessor(
		a.Blockdb,
		expireDuration, purgeDuration,
		cfg.MapBlockAccessorCfg.MaxItems)

	//color mapping
	a.Colormapping = colormapping.NewColorMapping()
	err = a.Colormapping.LoadDefaults()
	if err != nil {
		panic(err)
	}

	//load provided colors, if available
	info, err := os.Stat(filepath.Join(a.Config.ColorsTxtPath, "colors.txt"))
	if info != nil && err == nil {
		logrus.WithFields(logrus.Fields{"filename": "colors.txt"}).Info("Loading colors from filesystem")

		data, err := os.ReadFile(filepath.Join(a.Config.ColorsTxtPath, "colors.txt"))
		if err != nil {
			panic(err)
		}

		count, err := a.Colormapping.LoadBytes(data)
		if err != nil {
			panic(err)
		}

		logrus.WithFields(logrus.Fields{"count": count}).Info("Loaded custom colors")

	}

	//mapblock renderer
	a.Mapblockrenderer = mapblockrenderer.NewMapBlockRenderer(a.MapBlockAccessor, a.Colormapping)

	//mapserver database
	if a.Worldconfig[util.CONFIG_PSQL_MAPSERVER] != "" {
		a.Objectdb, err = postgresobjdb.New(a.Worldconfig[util.CONFIG_PSQL_MAPSERVER])
	} else {
		a.Objectdb, err = sqliteobjdb.New(filepath.Join(a.Config.DataPath, "mapserver.sqlite"))
	}

	if err != nil {
		panic(err)
	}

	//migrate object database
	err = a.Objectdb.Migrate()

	if err != nil {
		panic(err)
	}

	//create tiledb
	a.TileDB, err = tiledb.New(filepath.Join(a.Config.DataPath, "mapserver.tiles"))

	if err != nil {
		panic(err)
	}

	//settings
	a.Settings = settings.New(a.Objectdb)

	//setup tile renderer
	a.Tilerenderer = tilerenderer.NewTileRenderer(
		a.Mapblockrenderer,
		a.TileDB,
		a.Blockdb,
		a.Config.Layers,
	)

	//create media repo
	repo := make(map[string][]byte)

	if a.Config.EnableMediaRepository {
		mediasize, _ := util.ScanDir(repo, ".", []string{"mapserver.tiles", ".git"})
		fields := logrus.Fields{
			"count": len(repo),
			"bytes": mediasize,
		}
		logrus.WithFields(fields).Info("Created media repository")
	}

	a.MediaRepo = repo

	return &a
}
