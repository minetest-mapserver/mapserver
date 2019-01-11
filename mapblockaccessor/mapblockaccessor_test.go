package mapblockaccessor

import (
	"io"
	"io/ioutil"
	"mapserver/coords"
	"mapserver/db"
	"os"
	"testing"
	"github.com/sirupsen/logrus"
)

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

const testDatabase = "./testdata/map.sqlite"

func createTestDatabase(filename string) error {
	return copy(testDatabase, filename)
}

func GetTestDatabase() db.DBAccessor {
	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}

	createTestDatabase(tmpfile.Name())
	a, err := db.NewSqliteAccessor(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	return a
}

func TestSimpleAccess(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	a := GetTestDatabase()
	cache := NewMapBlockAccessor(a)
	mb, err := cache.GetMapBlock(coords.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if mb == nil {
		t.Fatal("Mapblock is nil")
	}
}
