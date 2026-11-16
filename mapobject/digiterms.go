package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type DigitermsBlock struct{}

func (blk *DigitermsBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "digiterm")
	o.Attributes["display_text"] = md["display_text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
