package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type JumpdriveBlock struct{}

func (this *JumpdriveBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "jumpdrive")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]

	return o
}
