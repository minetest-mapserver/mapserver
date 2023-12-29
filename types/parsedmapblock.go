package types

import (
	"github.com/minetest-go/mapparser"
)

type ParsedMapblock struct {
	Mapblock *mapparser.MapBlock
	Pos      *MapBlockCoords
}

func NewParsedMapblock(mb *mapparser.MapBlock, pos *MapBlockCoords) *ParsedMapblock {
	return &ParsedMapblock{Mapblock: mb, Pos: pos}
}
