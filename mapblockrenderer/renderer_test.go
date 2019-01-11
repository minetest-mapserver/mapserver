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

type JobData struct {
  pos1, pos2 coords.MapBlockCoords
  x,z int
}

func worker(r *MapBlockRenderer, jobs <-chan JobData) {
  for d := range jobs {
    img, _ := r.Render(d.pos1, d.pos2)

    if img != nil {
      f, _ := os.Create(fmt.Sprintf("../output/image_%d_%d.png", d.x, d.z))
      start := time.Now()
      png.Encode(f, img)
      f.Close()
      t := time.Now()
      elapsed := t.Sub(start)
      log.WithFields(logrus.Fields{"elapsed":elapsed}).Debug("Encoding completed")
    }
  }
}

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

  jobs := make(chan JobData, 100)
  go worker(&r, jobs)
  go worker(&r, jobs)
  go worker(&r, jobs)

  from := -1
  to := 1

  for x := from; x < to; x++ {
    for z := from; z < to; z++ {
      pos1 := coords.NewMapBlockCoords(x, 10, z)
      pos2 := coords.NewMapBlockCoords(x, -1, z)

      jobs <- JobData{pos1:pos1, pos2:pos2, x:x, z:z}
    }
  }

  close(jobs)


}
