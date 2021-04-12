package tilerenderer

import (
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db/sqlite"
	"mapserver/layer"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/testutils"
	"mapserver/tiledb"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func BenchmarkTileRender(b *testing.B) {
	logrus.SetLevel(logrus.PanicLevel)

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

	cache := mapblockaccessor.NewMapBlockAccessor(a, 500*time.Millisecond, 1000*time.Millisecond, 1000)
	c := colormapping.NewColorMapping()
	_, err = c.LoadVFSColors("colors/vanessa.txt")
	if err != nil {
		b.Fatal(err)
	}

	r := mapblockrenderer.NewMapBlockRenderer(cache, c)

	tiletmpdir, err := ioutil.TempDir("", "TestTileRenderTiles.*.sqlite")
	defer os.RemoveAll(tiletmpdir)

	tdb, _ := tiledb.New(tiletmpdir)

	layers := []*layer.Layer{
		&layer.Layer{
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
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		err := tr.Render(coord)
		if err != nil {
			panic(err)
		}

	}

}
