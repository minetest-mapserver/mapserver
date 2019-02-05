package tilerenderer

import (
	"bytes"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db/sqlite"
	"mapserver/layer"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"mapserver/testutils"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestTileRender(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	tmpfile, err := ioutil.TempFile("", "TestTileRender.*.sqlite")
	if err != nil {
		panic(err)
	}

	defer os.Remove(tmpfile.Name())
	testutils.CreateTestDatabase(tmpfile.Name())

	a, err := sqlite.New(tmpfile.Name())
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
		t.Fatal(err)
	}

	r := mapblockrenderer.NewMapBlockRenderer(cache, c)

	tiletmpfile, err := ioutil.TempFile("", "TestTileRenderTiles.*.sqlite")
	defer os.Remove(tiletmpfile.Name())

	tdb, _ := mapobjectdb.NewSqliteAccessor(tiletmpfile.Name())
	tdb.Migrate()

	layers := []layer.Layer{
		layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	tr := NewTileRenderer(r, tdb, a, layers)

	if tr == nil {
		panic("no renderer")
	}

	coord := coords.NewTileCoords(0, 0, 12, 0)
	data, err := tr.Render(coord, 2)
	if err != nil {
		panic(err)
	}

	if data == nil {
		panic("no data")
	}

	f, _ := os.Create("../output/0_0_12.png")
	bytes.NewReader(data).WriteTo(f)

	coord1 := coord.GetZoomedOutTile()

	data, err = tr.Render(coord1, 3)
	if err != nil {
		panic(err)
	}

	if data == nil {
		panic("no data")
	}

	f, _ = os.Create("../output/0_0_13.png")
	bytes.NewReader(data).WriteTo(f)

}
