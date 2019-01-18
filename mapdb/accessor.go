package mapdb

import (
	"mapserver/coords"
)

type MapData struct {
	//mapblock position
	MBPos   coords.MapBlockCoords
	//block position
	X, Y, Z int
	Type    string
	Data    string
	Mtime   int64
}

type SearchQuery struct {
	//block position (not mapblock)
	Pos1, Pos2 coords.MapBlockCoords
	Type string
}

type DBAccessor interface {
	Migrate() error
	GetMapData(q SearchQuery) ([]MapData, error)
	SetMapData(pos coords.MapBlockCoords, data []MapData) error
}
