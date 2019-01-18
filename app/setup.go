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
  "mapserver/layerconfig"

  "github.com/sirupsen/logrus"

  "errors"
)

func Setup(p params.ParamsType) (*App, error) {
  a := App{}
  a.Params = p

  //Parse world config

  a.Worldconfig = worldconfig.Parse(a.Params.Worlddir + "world.mt")
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

	a.Tiledb, err = tiledb.NewSqliteAccessor("tiles.sqlite")

	if err != nil {
		return nil, err
	}

  //migrate tile database

	err = a.Tiledb.Migrate()

	if err != nil {
		return nil, err
	}

  //setup tile renderer
  a.Tilerenderer = tilerenderer.NewTileRenderer(
    a.Mapblockrenderer,
    a.Tiledb,
    a.Blockdb,
    layerconfig.DefaultLayers,
  )

  return &a, nil
}
