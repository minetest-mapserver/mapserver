package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type SymbolBlock struct{}

func (this *SymbolBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "symbol")
	o.Attributes["minimap_text"] = md["minimap_text"]
	o.Attributes["minimap_symbol"] = md["minimap_symbol"]

	return o
}
