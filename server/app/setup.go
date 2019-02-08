package app

import (
	"mapserver/colormapping"
	"mapserver/db/sqlite"
	"mapserver/eventbus"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	sqliteobjdb "mapserver/mapobjectdb/sqlite"
	"mapserver/params"
	"mapserver/tilerenderer"
	"mapserver/worldconfig"
	"mapserver/settings"

	"github.com/sirupsen/logrus"

	"os"
	"io/ioutil"

	"errors"
)

func Setup(p params.ParamsType, cfg *Config) *App {
	a := App{}
	a.Params = p
	a.Config = cfg
	a.WebEventbus = eventbus.New()

	//Parse world config
	a.Worldconfig = worldconfig.Parse("world.mt")
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mapserver")

	var err error

	switch a.Worldconfig.Backend {
	case worldconfig.BACKEND_SQLITE3:
		a.Blockdb, err = sqlite.New("map.sqlite")
		if err != nil {
			panic(err)
		}
	default:
		panic(errors.New("map-backend not supported: " + a.Worldconfig.Backend))
	}

	//migrate block db
	err = a.Blockdb.Migrate()
	if err != nil {
		panic(err)
	}

	//mapblock accessor
	a.BlockAccessor = mapblockaccessor.NewMapBlockAccessor(a.Blockdb)

	//color mapping
	a.Colormapping = colormapping.NewColorMapping()

	//load default colors
	count, err := a.Colormapping.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		panic(err)
	}
	logrus.WithFields(logrus.Fields{"count": count}).Info("Loaded default colors")


	//load provided colors, if available
	info, err := os.Stat("colors.txt")
	if info != nil && err == nil {
		logrus.WithFields(logrus.Fields{"filename": "colors.txt"}).Info("Loading colors from filesystem")

		data, err := ioutil.ReadFile("colors.txt")
		if err != nil {
			panic(err)
		}

		count, err = a.Colormapping.LoadBytes(data)
		if err != nil {
			panic(err)
		}

		logrus.WithFields(logrus.Fields{"count": count}).Info("Loaded custom colors")

	}

	//mapblock renderer
	a.Mapblockrenderer = mapblockrenderer.NewMapBlockRenderer(a.BlockAccessor, a.Colormapping)

	//mapserver database
	if a.Worldconfig.MapObjectConnection != "" {
		//TODO: Psql connection

	} else {
		a.Objectdb, err = sqliteobjdb.New("mapserver.sqlite")
	}

	if err != nil {
		panic(err)
	}

	//migrate tile database
	err = a.Objectdb.Migrate()

	if err != nil {
		panic(err)
	}

	//settings
	a.Settings = settings.New(a.Objectdb)

	//setup tile renderer
	a.Tilerenderer = tilerenderer.NewTileRenderer(
		a.Mapblockrenderer,
		a.Objectdb,
		a.Blockdb,
		a.Config.Layers,
	)

	return &a
}
