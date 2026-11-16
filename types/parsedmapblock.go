package types

import (
	mtypes "github.com/luanti-golang/types"
)

type ParsedMapblock struct {
	Mapblock *mtypes.MapBlock
	Pos      *MapBlockCoords
}

func NewParsedMapblock(mb *mtypes.MapBlock, pos *MapBlockCoords) *ParsedMapblock {
	return &ParsedMapblock{Mapblock: mb, Pos: pos}
}
