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
	/**
	 * find old (pre-mapserver) mapblocks by lastpos
	 * used only on initial rendering
	 */
	FindLegacyBlocks(lastpos coords.MapBlockCoords, limit int) ([]Block, error)
	CountLegacyBlocks() (int, error)

	FindLatestBlocks(mintime int64, limit int) ([]Block, error)
	GetBlock(pos coords.MapBlockCoords) (*Block, error)
}
