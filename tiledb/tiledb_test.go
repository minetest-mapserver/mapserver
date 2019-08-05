package tiledb

import (
	"io/ioutil"
	"mapserver/coords"
	"os"
	"testing"
)

func TestTileDB(t *testing.T) {
	tmpfile, err := ioutil.TempDir("", "TestTileDB")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpfile)

	db, err := New(tmpfile)
	if err != nil {
		panic(err)
	}

	c := coords.NewTileCoords(0, 0, 1, 2)

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

	c2 := coords.NewTileCoords(1, 0, 1, 2)
	tile, err = db.GetTile(c2)
	if err != nil {
		panic(err)
	}

	if tile != nil {
		t.Error("tile exists")
	}

}
