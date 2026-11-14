package tiledb

import (
	"mapserver/util"
	"os"
	"testing"
)

func TestTileDB(t *testing.T) {
	tmpfile, err := os.MkdirTemp("", "TestTileDB")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpfile)

	db, err := New(tmpfile)
	if err != nil {
		panic(err)
	}

	c := util.NewTileCoords(0, 0, 1, 2)

	err = db.SetTile(c, []byte{1, 2, 3})
	if err != nil {
		panic(err)
	}

	tile, err := db.GetTile(c)
	if err != nil {
		panic(err)
	}

	if len(tile) != 3 {
		t.Error("wrong size")
	}

	c2 := util.NewTileCoords(1, 0, 1, 2)
	tile, err = db.GetTile(c2)
	if err != nil {
		panic(err)
	}

	if tile != nil {
		t.Error("tile exists")
	}

}
