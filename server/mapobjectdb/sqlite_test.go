package mapobjectdb

import (
	"fmt"
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

	pos := coords.NewTileCoords(0, 0, 13, 0)
	tile, err := db.GetTile(pos)
	if err != nil {
		panic(err)
	}

	if tile != nil {
		t.Fatal("non-empty tile found")
	}

	data := []byte{0x01, 0x02}
	tile2 := Tile{Pos: pos, Data: data}
	err = db.SetTile(&tile2)

	if err != nil {
		panic(err)
	}

	tile3, err := db.GetTile(pos)

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

func TestMapObjects(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMapObjects.*.sqlite")
	if err != nil {
		panic(err)
	}
	//defer os.Remove(tmpfile.Name())

	db, err := NewSqliteAccessor(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	attrs := make(map[string]string)
	attrs["X"] = "y"

	pos := coords.NewMapBlockCoords(0, 0, 0)

	o := MapObject{
		MBPos:      &pos,
		X:          1,
		Y:          2,
		Z:          3,
		Type:       "xy",
		Mtime:      1234,
		Attributes: attrs,
	}

	err = db.AddMapData(&o)
	if err != nil {
		panic(err)
	}

	q := SearchQuery{
		Pos1: pos,
		Pos2: pos,
		Type: "xy",
	}

	objs, err := db.GetMapData(q)

	if err != nil {
		panic(err)
	}

	for _, mo := range objs {
		fmt.Println(mo)
	}

}
