package app

import (
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/params"
	"mapserver/mapobjectdb"
	"mapserver/tilerenderer"
	"mapserver/worldconfig"

	"github.com/sirupsen/logrus"

	"errors"
)

func Setup(p params.ParamsType, cfg *Config) (*App, error) {
	a := App{}
	a.Params = p
	a.Config = cfg

	//Parse world config
	a.Worldconfig = worldconfig.Parse("world.mt")
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mapserver")

	if a.Worldconfig.Backend != worldconfig.BACKEND_SQLITE3 {
		return nil, errors.New("no supported backend found!")
	}

	//create db accessor
	var err error
	a.Blockdb, err = db.NewSqliteAccessor("map.sqlite")
	if err != nil {
		return nil, err
	}

	//migrate block db

	err = a.Blockdb.Migrate()
	if err != nil {
		return nil, err
	}

	//mapblock accessor
	a.BlockAccessor = mapblockaccessor.NewMapBlockAccessor(a.Blockdb)

	//color mapping

	a.Colormapping = colormapping.NewColorMapping()
	err = a.Colormapping.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		return nil, err
	}

	//mapblock renderer
	a.Mapblockrenderer = mapblockrenderer.NewMapBlockRenderer(a.BlockAccessor, a.Colormapping)

	//tile database

	a.Objectdb, err = mapobjectdb.NewSqliteAccessor("mapserver.sqlite")

	if err != nil {
		return nil, err
	}

	//migrate tile database

	err = a.Objectdb.Migrate()

	if err != nil {
		return nil, err
	}

	//setup tile renderer
	a.Tilerenderer = tilerenderer.NewTileRenderer(
		a.Mapblockrenderer,
		a.Objectdb,
		a.Blockdb,
		a.Config.Layers,
	)

	return &a, nil
}
