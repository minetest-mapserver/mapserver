package tiledb

import (
	"io/ioutil"
	"mapserver/coords"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TileDBTest.*.sqlite")
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
	tile, err := db.GetTile(0, pos)
	if err != nil {
		panic(err)
	}

	if tile != nil {
		t.Fatal("non-empty tile found")
	}

	data := []byte{0x01, 0x02}
	tile2 := Tile{LayerId: 0, Pos: pos, Data: data}
	err = db.SetTile(&tile2)

	if err != nil {
		panic(err)
	}

	tile3, err := db.GetTile(0, pos)

	if err != nil {
		panic(err)
	}

	if tile3 == nil {
		t.Fatal("no data returned")
	}

	if len(tile2.Data) != len(tile3.Data) {
		t.Fatal("inserted data does not match")
	}

	err = db.SetTile(&tile2)

	if err != nil {
		panic(err)
	}

}
