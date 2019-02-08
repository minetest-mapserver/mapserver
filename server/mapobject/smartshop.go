package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type SmartShopBlock struct{}

func (this *SmartShopBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "shop")
	o.Attributes["type"] = "smartshop"
	//TODO: 4 objects per coordinate

	//invMap := block.Metadata.GetInventoryMapAtPos(x, y, z)


	o.Attributes["index"] = md["index"]

	return o
}
