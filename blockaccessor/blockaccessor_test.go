package blockaccessor

import (
	"fmt"
	"io/ioutil"
	"mapserver/db/sqlite"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"os"
	"testing"
	"time"

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

	a, err := sqlite.New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	mba := mapblockaccessor.NewMapBlockAccessor(a, 500*time.Millisecond, 1000*time.Millisecond, 1000)

	if mba == nil {
		t.Fatal("Mapblockaccessor is nil")
	}

	ba := New(mba)

	if ba == nil {
		t.Fatal("blockaccessor is nil")
	}

	block, err := ba.GetBlock(0, 2, 0)

	if err != nil {
		panic(err)
	}

	if block == nil {
		t.Fatal("block is nil")
	}

	fmt.Println(block.Name)
}
