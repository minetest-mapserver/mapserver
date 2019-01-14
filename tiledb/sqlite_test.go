package tiledb

import (
	"io/ioutil"
	"os"
	"testing"
	"mapserver/coords"
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

	pos := coords.NewTileCoords(0,0,13)
	_, err = db.GetTile(0, pos)
	if err != nil {
		panic(err)
	}


}
