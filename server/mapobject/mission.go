package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type MissionBlock struct{}

func (this *MissionBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "mission")
	o.Attributes["name"] = md["name"]
	o.Attributes["time"] = md["time"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["description"] = md["description"]
	o.Attributes["successcount"] = md["successcount"]
	o.Attributes["failcount"] = md["failcount"]

	return o
}
