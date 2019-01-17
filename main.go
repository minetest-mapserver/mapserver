package main

import (
	"flag"
	"fmt"
	"mapserver/initialrenderer"
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/params"
	"mapserver/worldconfig"
	"github.com/sirupsen/logrus"
	"mapserver/tilerenderer"
	"mapserver/tiledb"
	"mapserver/layerconfig"

)

const (
	Version = "2.0-DEV"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	p := params.Parse()

	if p.Help {
		flag.PrintDefaults()
		return
	}

	if p.Version {
		fmt.Print("Mapserver version: ")
		fmt.Println(Version)
		return
	}

	worldcfg := worldconfig.Parse(p.Worlddir + "world.mt")
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mapserver")

	if worldcfg.Backend != worldconfig.BACKEND_SQLITE3 {
		panic("no sqlite3 backend found!")
	}

	a, err := db.NewSqliteAccessor("map.sqlite")
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	cache := mapblockaccessor.NewMapBlockAccessor(a)
	c := colormapping.NewColorMapping()
	err = c.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		panic(err)
	}

	r := mapblockrenderer.NewMapBlockRenderer(cache, c)

	tdb, err := tiledb.NewSqliteAccessor("tiles.sqlite")

	if err != nil {
		panic(err)
	}

	err = tdb.Migrate()

	if err != nil {
		panic(err)
	}

	tr := tilerenderer.NewTileRenderer(r, tdb, a, layerconfig.DefaultLayers)

	initialrenderer.Render(tr, layerconfig.DefaultLayers)

}
