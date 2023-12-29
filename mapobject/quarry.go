package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type QuarryBlock struct{}

func (this *QuarryBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
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
