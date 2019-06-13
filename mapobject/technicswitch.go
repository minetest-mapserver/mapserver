package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type TechnicSwitchBlock struct{}

func (this *TechnicSwitchBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "technicswitch")
	o.Attributes["active"] = md["active"]
	o.Attributes["channel"] = md["channel"]
	o.Attributes["supply"] = md["supply"]
	o.Attributes["demand"] = md["demand"]

	return o
}
