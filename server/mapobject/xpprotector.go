package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type XPProtectorBlock struct{}

func (this *XPProtectorBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "xpprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["xpthreshold"] = md["xpthreshold"]

	return o
}
