package mapblockrenderer

import (
	"fmt"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db/sqlite"
	"mapserver/layer"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestSimpleRender(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	layers := []*layer.Layer{
		&layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
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
		t.Fatal(err)
	}

	r := NewMapBlockRenderer(cache, c)
	os.Mkdir("../test-output", 0755)

	results := make(chan JobResult, 100)
	jobs := make(chan JobData, 100)

	go Worker(r, jobs, results)
	go Worker(r, jobs, results)
	go Worker(r, jobs, results)

	go func() {
		for result := range results {
			if result.Data.Len() == 0 {
				continue
			}

			tc := coords.GetTileCoordsFromMapBlock(result.Job.Pos1, layers)
			f, _ := os.Create(fmt.Sprintf("../test-output/image_%d_%d.png", tc.X, tc.Y))
			result.Data.WriteTo(f)
			f.Close()
		}
	}()

	from := -10
	to := 10

	for x := from; x < to; x++ {
		for z := from; z < to; z++ {
			pos1 := coords.NewMapBlockCoords(x, 10, z)
			pos2 := coords.NewMapBlockCoords(x, -1, z)

			jobs <- JobData{Pos1: pos1, Pos2: pos2}
		}
	}

	close(jobs)

}
