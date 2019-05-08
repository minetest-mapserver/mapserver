package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type Locator struct{}

func (this *Locator) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)
	nodename := block.GetNodeName(x, y, z)

	var level = "1"
	switch nodename {
	case "locator:beacon_2":
		level = "2"
	case "locator:beacon_3":
		level = "3"
	}

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "locator")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["name"] = md["name"]
	o.Attributes["active"] = md["active"]
	o.Attributes["level"] = level

	return o
}
