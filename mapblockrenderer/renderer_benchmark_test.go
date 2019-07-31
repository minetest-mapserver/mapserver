package mapblockrenderer

import (
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db/sqlite"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type callback func()

func createRenderer(b *testing.B) (*MapBlockRenderer, callback) {
	logrus.SetLevel(logrus.PanicLevel)

	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		os.Remove(tmpfile.Name())
	}

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
	_, err = c.LoadVFSColors(false, "/colors/vanessa.txt")
	if err != nil {
		b.Fatal(err)
	}

	r := NewMapBlockRenderer(cache, c)
	b.ResetTimer()

	return r, cleanup
}


func BenchmarkRenderEmptySingle(b *testing.B) {
	r, cleanup := createRenderer(b)
	defer cleanup()

	for n := 0; n < b.N; n++ {

		pos1 := coords.NewMapBlockCoords(10, 0, 10)
		pos2 := coords.NewMapBlockCoords(10, 0, 10)

		_, err := r.Render(pos1, pos2)

		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkRenderSingle(b *testing.B) {
	r, cleanup := createRenderer(b)
	defer cleanup()

	for n := 0; n < b.N; n++ {

		pos1 := coords.NewMapBlockCoords(0, 0, 0)
		pos2 := coords.NewMapBlockCoords(0, 0, 0)

		_, err := r.Render(pos1, pos2)

		if err != nil {
			panic(err)
		}
	}
}


func BenchmarkRenderStride(b *testing.B) {
	r, cleanup := createRenderer(b)
	defer cleanup()

	for n := 0; n < b.N; n++ {

		pos1 := coords.NewMapBlockCoords(0, 10, 0)
		pos2 := coords.NewMapBlockCoords(0, -1, 0)

		_, err := r.Render(pos1, pos2)

		if err != nil {
			panic(err)
		}
	}
}


func BenchmarkRenderBigStride(b *testing.B) {
	r, cleanup := createRenderer(b)
	defer cleanup()

	for n := 0; n < b.N; n++ {

		pos1 := coords.NewMapBlockCoords(0, 1000, 0)
		pos2 := coords.NewMapBlockCoords(0, -1000, 0)

		_, err := r.Render(pos1, pos2)

		if err != nil {
			panic(err)
		}
	}
}

