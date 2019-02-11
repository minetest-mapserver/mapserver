package db

import (
	"mapserver/coords"
)

type Block struct {
	Pos   *coords.MapBlockCoords
	Data  []byte
	Mtime int64
}

type InitialBlocksResult struct {
	List    []*Block
	HasMore bool
}

type DBAccessor interface {
	Migrate() error

	FindBlocksByMtime(gtmtime int64, limit int) ([]*Block, error)
	FindNextInitialBlocks(lastpos *coords.MapBlockCoords, limit int) (*InitialBlocksResult, error)

	CountBlocks(frommtime, tomtime int64) (int, error)
	GetBlock(pos *coords.MapBlockCoords) (*Block, error)
}
