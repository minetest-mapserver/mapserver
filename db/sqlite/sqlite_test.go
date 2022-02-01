package sqlite

import (
	"io/ioutil"
	"mapserver/coords"
	"mapserver/testutils"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func TestMigrateEmpty(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrateEmpty.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}
}

func TestMigrate(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}
}

func TestMigrateAndQuery(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrateAndQuery.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	block, err := a.GetBlock(coords.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if block == nil {
		t.Fatal("no data")
	}

}

func TestMigrateAndQueryCount(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	count, err := a.CountBlocks()
	if err != nil {
		panic(err)
	}

	if count <= 0 {
		t.Fatal("zero count")
	}
}

func TestMigrateAndQueryTimestamp(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	count, err := a.GetTimestamp()
	if err != nil {
		panic(err)
	}

	if count <= 0 {
		t.Fatal("zero count")
	}
}
