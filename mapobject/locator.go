package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type Locator struct{}

func (this *Locator) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)
	nodename := block.GetNodeName(x, y, z)

	var level = "1"
	switch nodename {
	case "locator:beacon_2":
		level = "2"
	case "locator:beacon_3":
		level = "3"
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "locator")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["name"] = md["name"]
	o.Attributes["active"] = md["active"]
	o.Attributes["level"] = level

	return o
}
