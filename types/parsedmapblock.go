package types

import (
	"mapserver/coords"

	"github.com/minetest-go/mapparser"
)

type ParsedMapblock struct {
	Mapblock *mapparser.MapBlock
	Pos      *coords.MapBlockCoords
}

func NewParsedMapblock(mb *mapparser.MapBlock, pos *coords.MapBlockCoords) *ParsedMapblock {
	return &ParsedMapblock{Mapblock: mb, Pos: pos}
}
