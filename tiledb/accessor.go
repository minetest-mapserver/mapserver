package tiledb

import (
	"mapserver/coords"
)

type Tile struct {
	Pos     coords.TileCoords
	LayerId int
	Data    []byte
	Mtime   int64
}

type DBAccessor interface {
	Migrate() error
	GetTile(pos coords.TileCoords) (*Tile, error)
	SetTile(pos coords.TileCoords, tile *Tile) error
}
