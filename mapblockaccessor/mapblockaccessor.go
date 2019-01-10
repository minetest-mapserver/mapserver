package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockparser"
)

type MapBlockAccessor struct {
	accessor db.DBAccessor
}

func NewMapBlockAccessor(accessor db.DBAccessor) *MapBlockAccessor {
	return &MapBlockAccessor{accessor: accessor}
}

func (a *MapBlockAccessor) Update(pos coords.MapBlockCoords, mb *mapblockparser.MapBlock) {
	//TODO: cache
}

func (a *MapBlockAccessor) GetMapBlock(pos coords.MapBlockCoords) (*mapblockparser.MapBlock, error) {
	block, err := a.accessor.GetBlock(pos)
	if err != nil {
		return nil, err
	}

	mapblock, err := mapblockparser.Parse(block.Data)
	if err != nil {
		return nil, err
	}

	return mapblock, nil
}
