package mapobjectdb

import (
	"mapserver/coords"
)

type MapObject struct {
	//mapblock position
	MBPos coords.MapBlockCoords

	//block position
	X, Y, Z int

	Type  string
	Data  string
	Mtime int64
}

type SearchQuery struct {
	//block position (not mapblock)
	Pos1, Pos2 coords.MapBlockCoords
	Type       string
}

type DBAccessor interface {
	Migrate() error
	GetMapData(q SearchQuery) ([]MapObject, error)
	RemoveMapData(pos coords.MapBlockCoords) error
	AddMapData(data MapObject) error
}
