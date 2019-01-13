package mapblockrenderer

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type JobData struct {
	pos1, pos2 coords.MapBlockCoords
	x, z       int
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
			logrus.WithFields(logrus.Fields{"elapsed": elapsed}).Debug("Encoding completed")
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

	from := -10
	to := 10

	for x := from; x < to; x++ {
		for z := from; z < to; z++ {
			pos1 := coords.NewMapBlockCoords(x, 10, z)
			pos2 := coords.NewMapBlockCoords(x, -1, z)

			jobs <- JobData{pos1: pos1, pos2: pos2, x: x, z: z}
		}
	}

	close(jobs)

}
