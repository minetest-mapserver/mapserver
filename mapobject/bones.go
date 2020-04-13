package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
	"strconv"
)

type BonesBlock struct{}

func (this *BonesBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	invMap := block.Metadata.GetInventoryMapAtPos(x, y, z)
	mainInv := invMap["main"]

	if mainInv == nil {
		return nil
	}

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "bones")
	o.Attributes["time"] = md["time"]

	if _, ok := md["owner"]; ok {
		o.Attributes["owner"] = md["owner"]
	} else if _, ok := md["_owner"]; ok {
		o.Attributes["owner"] = md["_owner"]
	} else {
		o.Attributes["owner"] = "unknown"
	}

	o.Attributes["info"] = md["infotext"]

	itemCount := 0
	for _, item := range mainInv.Items {
		itemCount += item.Count
	}

	o.Attributes["item_count"] = strconv.Itoa(itemCount)

	return o
}
