package mapobjectdb

import (
	"mapserver/coords"
)

func (db *Sqlite3Accessor) GetMapData(q SearchQuery) ([]MapObject, error) {
	return nil, nil
}

func (db *Sqlite3Accessor) RemoveMapData(pos coords.MapBlockCoords) error {
	return nil
}

func (db *Sqlite3Accessor) AddMapData(data MapObject) error {
	return nil
}
