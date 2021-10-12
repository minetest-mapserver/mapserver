package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type TechnicSwitchBlock struct{}

func (this *TechnicSwitchBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicswitch")
	o.Attributes["active"] = md["active"]
	o.Attributes["channel"] = md["channel"]
	o.Attributes["supply"] = md["supply"]
	o.Attributes["demand"] = md["demand"]

	return o
}
