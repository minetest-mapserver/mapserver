package tiledb

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMigrateEmpty.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	db, err := NewSqliteAccessor(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}
}
