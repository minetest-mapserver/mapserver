package db

import (
	"mapserver/settings"
	"mapserver/types"
)

type Block struct {
	Pos   *types.MapBlockCoords
	Data  []byte
	Mtime int64
}

type InitialBlocksResult struct {
	List            []*Block
	UnfilteredCount int
	HasMore         bool
	Progress        float64
	LastMtime       int64
}

type DBAccessor interface {
	Migrate() error

	GetTimestamp() (int64, error)
	FindBlocksByMtime(gtmtime int64, limit int) ([]*Block, error)
	FindNextInitialBlocks(s settings.Settings, layers []*types.Layer, limit int) (*InitialBlocksResult, error)
	GetBlock(pos *types.MapBlockCoords) (*Block, error)
}
