package mapblockrenderer

import (
  "os"
	"io/ioutil"
	"mapserver/coords"
	"testing"
	"mapserver/testutils"
	"mapserver/db"
  "mapserver/colormapping"
  "mapserver/mapblockaccessor"
  "image/png"
)

func TestSimpleRender(t *testing.T) {
  tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
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

  r := NewMapBlockRenderer(cache, c)
  img, _ := r.Render(coords.NewMapBlockCoords(0, 10, 0), coords.NewMapBlockCoords(0, -1, 0))

  f, _ := os.Create("image.png")
  png.Encode(f, img)
  f.Close()
}
