package mapobject

import (
	"mapserver/coords"
	"mapserver/mapblockparser"
	"mapserver/mapdb"
)

type ClearMapData struct {
	db mapdb.DBAccessor
}

func (this *ClearMapData) OnParsedMapBlock(block *mapblockparser.MapBlock, pos coords.MapBlockCoords) {
	err := this.db.RemoveMapData(pos)
	if err != nil {
		panic(err)
	}
}
