package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type Phonograph struct{}

func (this *Phonograph) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if _, ok := md["song_title"]; !ok {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "phonograph")
	o.Attributes["song_title"] = md["song_title"]
	o.Attributes["song_artist"] = md["song_artist"]

	return o
}
