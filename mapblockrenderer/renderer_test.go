package mapblockrenderer

import (
  "os"
	"io/ioutil"
	"mapserver/coords"
	"testing"
  "fmt"
	"mapserver/testutils"
	"mapserver/db"
  "mapserver/colormapping"
  "mapserver/mapblockaccessor"
  "image/png"
  "github.com/sirupsen/logrus"
  "time"
)

func TestSimpleRender(t *testing.T) {
  logrus.SetLevel(logrus.InfoLevel)

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
  os.Mkdir("../output", 0755)

  for x := -3; x < 3; x++ {
    for z := -3; z < 3; z++ {
      img, _ := r.Render(coords.NewMapBlockCoords(x, 10, z), coords.NewMapBlockCoords(x, -1, z))

      if img != nil {
        f, _ := os.Create(fmt.Sprintf("../output/image_%d_%d.png", x, z))
        start := time.Now()
        png.Encode(f, img)
        f.Close()
        t := time.Now()
        elapsed := t.Sub(start)
        log.WithFields(logrus.Fields{"elapsed":elapsed}).Debug("Encoding completed")

      }
    }
  }

}
