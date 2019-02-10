package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
)

func (a *MapBlockAccessor) GetMapBlock(pos *coords.MapBlockCoords) (*mapblockparser.MapBlock, error) {
	key := getKey(pos)

	cachedblock, found := a.c.Get(key)
	if found {
		return cachedblock.(*mapblockparser.MapBlock), nil
	}

	block, err := a.accessor.GetBlock(pos)
	if err != nil {
		return nil, err
	}

	if block == nil {
		return nil, nil
	}

	mapblock, err := mapblockparser.Parse(block.Data, block.Mtime, pos)
	if err != nil {
		return nil, err
	}

	a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, mapblock)

	a.c.Set(key, mapblock, cache.DefaultExpiration)

	return mapblock, nil
}
