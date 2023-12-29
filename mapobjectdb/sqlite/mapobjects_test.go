package sqlite

import (
	"fmt"
	"io/ioutil"
	"mapserver/mapobjectdb"
	"mapserver/types"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TileDBTest.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	db, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
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

	db, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	attrs := make(map[string]string)
	attrs["X"] = "y"

	pos := types.NewMapBlockCoords(0, 0, 0)

	o := mapobjectdb.MapObject{
		MBPos:      pos,
		X:          15,
		Y:          15,
		Z:          15,
		Type:       "xy",
		Mtime:      1234,
		Attributes: attrs,
	}

	err = db.AddMapData(&o)
	if err != nil {
		panic(err)
	}

	limit := 1000
	q := mapobjectdb.SearchQuery{
		Pos1:  pos,
		Pos2:  pos,
		Type:  "xy",
		Limit: &limit,
	}

	objs, err := db.GetMapData(&q)

	if err != nil {
		panic(err)
	}

	for _, mo := range objs {
		fmt.Println(mo)
	}

}

func TestMapObjectsQueryWithAttribute(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMapObjects.*.sqlite")
	if err != nil {
		panic(err)
	}
	//defer os.Remove(tmpfile.Name())

	db, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	attrs := make(map[string]string)
	attrs["X"] = "y"

	pos := types.NewMapBlockCoords(0, 0, 0)

	o := mapobjectdb.MapObject{
		MBPos:      pos,
		X:          15,
		Y:          15,
		Z:          15,
		Type:       "xy",
		Mtime:      1234,
		Attributes: attrs,
	}

	err = db.AddMapData(&o)
	if err != nil {
		panic(err)
	}

	limit := 1000
	q := mapobjectdb.SearchQuery{
		Pos1: pos,
		Pos2: pos,
		Type: "xy",
		AttributeLike: &mapobjectdb.SearchAttributeLike{
			Key:   "X",
			Value: "%y",
		},
		Limit: &limit,
	}

	objs, err := db.GetMapData(&q)

	if err != nil {
		panic(err)
	}

	for _, mo := range objs {
		fmt.Println(mo)
	}

	if len(objs) != 1 {
		panic("length mismatch")
	}
}

func TestMapObjectsQueryWithAttributeIgnoreCase(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestMapObjects.*.sqlite")
	if err != nil {
		panic(err)
	}
	//defer os.Remove(tmpfile.Name())

	db, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	attrs := make(map[string]string)
	attrs["X"] = "ABC"

	pos := types.NewMapBlockCoords(0, 0, 0)

	o := mapobjectdb.MapObject{
		MBPos:      pos,
		X:          15,
		Y:          15,
		Z:          15,
		Type:       "xy",
		Mtime:      1234,
		Attributes: attrs,
	}

	err = db.AddMapData(&o)
	if err != nil {
		panic(err)
	}

	limit := 1000
	q := mapobjectdb.SearchQuery{
		Pos1: pos,
		Pos2: pos,
		Type: "xy",
		AttributeLike: &mapobjectdb.SearchAttributeLike{
			Key:   "X",
			Value: "%bc",
		},
		Limit: &limit,
	}

	objs, err := db.GetMapData(&q)

	if err != nil {
		panic(err)
	}

	for _, mo := range objs {
		fmt.Println(mo)
	}

	if len(objs) != 1 {
		panic("length mismatch")
	}
}
