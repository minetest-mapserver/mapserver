package mapblockaccessor

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockparser"
	"time"
)

type MapBlockAccessor struct {
	accessor db.DBAccessor
	c        *cache.Cache
	listeners []MapBlockListener
}

type MapBlockListener interface {
	OnParsedMapBlock(block *mapblockparser.MapBlock)
}

func getKey(pos coords.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor) *MapBlockAccessor {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return &MapBlockAccessor{accessor: accessor, c: c}
}

func (a *MapBlockAccessor) AddListener(l MapBlockListener) {
	a.listeners = append(a.listeners, l)
}

func (a *MapBlockAccessor) Update(pos coords.MapBlockCoords, mb *mapblockparser.MapBlock) {
	key := getKey(pos)
	a.c.Set(key, mb, cache.DefaultExpiration)
}

func (a *MapBlockAccessor) FindLatestMapBlocks(mintime int64, limit int) ([]*mapblockparser.MapBlock, error) {
	blocks, err := a.accessor.FindLatestBlocks(mintime, limit)

	if err != nil {
		return nil, err
	}

	mblist := make([]*mapblockparser.MapBlock, 0)

	for _, block := range blocks {
		key := getKey(block.Pos)

		mapblock, err := mapblockparser.Parse(block.Data, block.Mtime)
		if err != nil {
			return nil, err
		}

		for _, listener := range(a.listeners) {
			listener.OnParsedMapBlock(mapblock)
		}

		a.c.Set(key, mapblock, cache.DefaultExpiration)
		mblist = append(mblist, mapblock)
	}

	return mblist, nil
}

func (a *MapBlockAccessor) GetMapBlock(pos coords.MapBlockCoords) (*mapblockparser.MapBlock, error) {
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

	mapblock, err := mapblockparser.Parse(block.Data, block.Mtime)
	if err != nil {
		return nil, err
	}

	for _, listener := range(a.listeners) {
		listener.OnParsedMapBlock(mapblock)
	}

	a.c.Set(key, mapblock, cache.DefaultExpiration)

	return mapblock, nil
}
