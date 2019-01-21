package mapblockrenderer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"mapserver/layer"
	"os"
	"testing"
)

func TestSimpleRender(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	layers := []layer.Layer{
		layer.Layer{
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
			f, _ := os.Create(fmt.Sprintf("../output/image_%d_%d.png", tc.X, tc.Y))
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
