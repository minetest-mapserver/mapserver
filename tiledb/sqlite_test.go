package tiledb

import (
	"io/ioutil"
	"mapserver/coords"
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

	pos := coords.NewTileCoords(0, 0, 13)
	_, err = db.GetTile(0, pos)
	if err != nil {
		panic(err)
	}

	data := []byte{}
	tile := Tile{LayerId: 0, Pos: pos, Data: data}
	err = db.SetTile(&tile)

	if err != nil {
		panic(err)
	}

}
