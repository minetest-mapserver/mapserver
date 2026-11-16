package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type Phonograph struct{}

func (phg *Phonograph) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	if _, ok := md["song_title"]; !ok {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "phonograph")
	o.Attributes["song_title"] = md["song_title"]
	o.Attributes["song_artist"] = md["song_artist"]

	return o
}
