package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type Locator struct{}

func (l *Locator) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)
	nodename := block.GetNodeName(mtypes.NewPos(x, y, z))

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
