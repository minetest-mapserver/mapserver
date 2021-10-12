package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type XPProtectorBlock struct{}

func (this *XPProtectorBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "xpprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["xpthreshold"] = md["xpthreshold"]

	return o
}
