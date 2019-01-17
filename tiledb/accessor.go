package tiledb

import (
	"mapserver/coords"
)

type Tile struct {
	Pos     coords.TileCoords
	Data    []byte
	Mtime   int64
}

type DBAccessor interface {
	Migrate() error
	GetTile(pos coords.TileCoords) (*Tile, error)
	SetTile(tile *Tile) error
}
