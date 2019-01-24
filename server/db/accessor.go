package db

import (
	"mapserver/coords"
)

type Block struct {
	Pos   coords.MapBlockCoords
	Data  []byte
	Mtime int64
}

type DBAccessor interface {
	Migrate() error
	FindBlocks(lastpos coords.MapBlockCoords, lastmtime int64, limit int) ([]Block, error)
	CountBlocks(frommtime, tomtime int64) (int, error)
	GetBlock(pos coords.MapBlockCoords) (*Block, error)
}
