package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
	"fmt"
)

type SymbolBlock struct{}

func (this *SymbolBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	fmt.Printf("Symbol at %d %d %d", x, y, z)
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "symbol")
	o.Attributes["minimap_text"] = md["minimap_text"]
	o.Attributes["minimap_symbol"] = md["minimap_symbol"]

	return o
}
