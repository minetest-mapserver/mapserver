package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
	"strconv"
)

type BonesBlock struct{}

func (this *BonesBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["owner"] == "" {
		return nil
	}

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "bones")
	o.Attributes["time"] = md["time"]
	o.Attributes["owner"] = md["owner"]

	invMap := block.Metadata.GetInventoryMapAtPos(x, y, z)
	mainInv := invMap["main"]

	itemCount := 0
	for _, item := range mainInv.Items {
		itemCount += item.Count
	}

	o.Attributes["item_count"] = strconv.Itoa(itemCount)

	return o
}
