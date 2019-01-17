package tilerenderer

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/db"
	"mapserver/mapblockaccessor"
  "mapserver/testutils"
  "mapserver/mapblockrenderer"
  "mapserver/coords"
	"os"
	"testing"
)

func TestTileRender(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	tmpfile, err := ioutil.TempFile("", "TestTileRender.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())
	testutils.CreateTestDatabase(tmpfile.Name())

	a, err := db.NewSqliteAccessor(tmpfile.Name())
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

  tr := NewTileRenderer(r)

  if tr == nil {
    panic("no renderer")
  }

  coord := coords.NewTileCoords(0,0,12,0)

  data, err := tr.Render(coord)

  if err != nil {
    panic(err)
  }

  if data == nil {
    panic("no data")
  }
}
