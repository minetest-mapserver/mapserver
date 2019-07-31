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

func BenchmarkRender(b *testing.B) {
	logrus.SetLevel(logrus.ErrorLevel)

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
	_, err = c.LoadVFSColors(false, "/colors/vanessa.txt")
	if err != nil {
		b.Fatal(err)
	}

	r := NewMapBlockRenderer(cache, c)
	os.Mkdir("../test-output", 0755)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		pos1 := coords.NewMapBlockCoords(0, 10, 0)
		pos2 := coords.NewMapBlockCoords(0, -1, 0)

		_, err := r.Render(pos1, pos2)

		if err != nil {
			panic(err)
		}
	}
}

