package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type SignBlock struct {
	Material string
}

func (blk *SignBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "sign")
	o.Attributes["display_text"] = md["text"]
	o.Attributes["material"] = blk.Material

	return o
}
