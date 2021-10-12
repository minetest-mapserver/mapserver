package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type QuarryBlock struct{}

func (this *QuarryBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["owner"] == "" {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicquarry")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["dug"] = md["dug"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
