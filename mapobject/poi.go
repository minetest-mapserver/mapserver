package mapobject

import (
	"mapserver/coords"
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type POI struct {
	db mapobjectdb.DBAccessor
}

func (this *POI) OnParsedMapBlock(block *mapblockparser.MapBlock, pos coords.MapBlockCoords) {
	var found bool
	for _, v := range block.BlockMapping {
		if v == "mapserver:poi" {
			found = true
			break
		}
	}

	if !found {
		return
	}

	panic("OK") //XXX
}
