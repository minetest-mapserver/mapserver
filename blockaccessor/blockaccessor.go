package blockaccessor

import (
	"mapserver/coords"
	"mapserver/mapblockaccessor"
)

func New(mba *mapblockaccessor.MapBlockAccessor) *BlockAccessor {
	return &BlockAccessor{mba: mba}
}

type BlockAccessor struct {
	mba *mapblockaccessor.MapBlockAccessor
}

type Block struct {
	Name string
	//TODO: param1, param2
}

func (this *BlockAccessor) GetBlock(x, y, z int) (*Block, error) {

	mbc := coords.NewMapBlockCoordsFromBlock(x, y, z)
	mapblock, err := this.mba.GetMapBlock(mbc)

	if err != nil {
		return nil, err
	}

	if mapblock == nil {
		return nil, nil
	}

	relx := x % 16
	rely := y % 16
	relz := z % 16

	block := Block{
		Name: mapblock.GetNodeName(relx, rely, relz),
	}

	return &block, nil
}
