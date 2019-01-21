package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type ClearMapData struct {
	db mapobjectdb.DBAccessor
}

func (this *ClearMapData) OnParsedMapBlock(block *mapblockparser.MapBlock) {
	err := this.db.RemoveMapData(block.Pos)
	if err != nil {
		panic(err)
	}
}
