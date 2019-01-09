package db

import (
	"mapserver/coords"
)

type Block struct {
	pos coords.MapBlockCoords
	data             []byte
	mtime            int64
}

type DBAccessor interface {
	Migrate() error
	FindLatestBlocks(mintime int64, limit int) ([]Block, error)
	FindBlocks(pos1, pos2 coords.MapBlockCoords) ([]Block, error)
	CountBlocks(pos1, pos2 coords.MapBlockCoords) (int, error)
}
