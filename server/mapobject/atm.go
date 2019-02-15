package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type ATM struct{}

func (this *ATM) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	return mapobjectdb.NewMapObject(block.Pos, x, y, z, "atm")
}
