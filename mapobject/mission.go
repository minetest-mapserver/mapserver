package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type MissionBlock struct{}

func (this *MissionBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["hidden"] == "1" {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "mission")
	o.Attributes["name"] = md["name"]
	o.Attributes["time"] = md["time"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["description"] = md["description"]
	o.Attributes["successcount"] = md["successcount"]
	o.Attributes["failcount"] = md["failcount"]

	return o
}
