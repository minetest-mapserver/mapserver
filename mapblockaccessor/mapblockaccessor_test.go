package mapblockaccessor

import (
	"os"
	"io/ioutil"
	"mapserver/coords"
	"testing"
	"mapserver/testutils"
	"mapserver/db"
	"github.com/sirupsen/logrus"
)

func TestSimpleAccess(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

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

	cache := NewMapBlockAccessor(a)
	mb, err := cache.GetMapBlock(coords.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if mb == nil {
		t.Fatal("Mapblock is nil")
	}
}
