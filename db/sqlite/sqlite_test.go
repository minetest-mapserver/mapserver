package sqlite

import (
	"mapserver/testutils"
	"mapserver/types"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func TestMigrateEmpty(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateEmpty.*.sqlite")
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

func TestMigrateEmptyLegacy(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateEmpty.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyLegacyDatabase(tmpfile.Name())
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
	tmpfile, err := os.CreateTemp("", "TestMigrate.*.sqlite")
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
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQuery.*.sqlite")
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

	block, err := a.GetBlock(types.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if block == nil {
		t.Fatal("no data")
	}

}

func TestMigrateAndQueryCount(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
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

func TestMigrateAndQueryCountLegacy(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabaseLegacy(tmpfile.Name())
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
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
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
