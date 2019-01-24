package mapobjectdb

import (
	"mapserver/coords"
	"time"
)

/*
sqlite perf: https://stackoverflow.com/questions/1711631/improve-insert-per-second-performance-of-sqlite?rq=1
PRAGMA synchronous = OFF
PRAGMA journal_mode = MEMORY
*/

type Tile struct {
	Pos   *coords.TileCoords
	Data  []byte
	Mtime int64
}

type MapObject struct {
	//mapblock position
	MBPos *coords.MapBlockCoords

	//block position
	X, Y, Z int

	Type       string
	Mtime      int64
	Attributes map[string]string
}

func NewMapObject(MBPos *coords.MapBlockCoords, x int, y int, z int, _type string) *MapObject {
	o := MapObject{
		MBPos:      MBPos,
		Type:       _type,
		X:          MBPos.X + x,
		Y:          MBPos.Y + y,
		Z:          MBPos.Z + z,
		Mtime:      time.Now().Unix(),
		Attributes: make(map[string]string),
	}

	return &o
}

type SearchQuery struct {
	//block position (not mapblock)
	Pos1, Pos2 coords.MapBlockCoords
	Type       string
}

type DBAccessor interface {
	Migrate() error

	//Generic map objects (poi, etc)
	GetMapData(q SearchQuery) ([]MapObject, error)
	RemoveMapData(pos *coords.MapBlockCoords) error
	AddMapData(data *MapObject) error

	//tile data
	GetTile(pos *coords.TileCoords) (*Tile, error)
	SetTile(tile *Tile) error
	RemoveTile(pos *coords.TileCoords) error
}
