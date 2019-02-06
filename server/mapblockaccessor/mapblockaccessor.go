package mapblockaccessor

import (
	"fmt"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/layer"
	"mapserver/mapblockparser"

	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type MapBlockAccessor struct {
	accessor db.DBAccessor
	c        *cache.Cache
	Eventbus *eventbus.Eventbus
}

func getKey(pos *coords.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor) *MapBlockAccessor {
	c := cache.New(500*time.Millisecond, 1000*time.Millisecond)

	return &MapBlockAccessor{
		accessor: accessor,
		c:        c,
		Eventbus: eventbus.New(),
	}
}
func (a *MapBlockAccessor) Update(pos *coords.MapBlockCoords, mb *mapblockparser.MapBlock) {
	key := getKey(pos)
	a.c.Set(key, mb, cache.DefaultExpiration)
}

type FindMapBlocksResult struct {
	HasMore         bool
	LastPos         *coords.MapBlockCoords
	LastMtime       int64
	List            []*mapblockparser.MapBlock
	UnfilteredCount int
}

func (a *MapBlockAccessor) FindMapBlocksByMtime(lastmtime int64, limit int, layerfilter []layer.Layer) (*FindMapBlocksResult, error) {

	fields := logrus.Fields{
		"lastmtime": lastmtime,
		"limit":     limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByMtime")

	blocks, err := a.accessor.FindBlocksByMtime(lastmtime, limit)

	if err != nil {
		return nil, err
	}

	result := FindMapBlocksResult{}

	mblist := make([]*mapblockparser.MapBlock, 0)
	var newlastpos *coords.MapBlockCoords
	result.HasMore = len(blocks) == limit
	result.UnfilteredCount = len(blocks)

	for _, block := range blocks {
		newlastpos = block.Pos
		if result.LastMtime < block.Mtime {
			result.LastMtime = block.Mtime
		}

		inLayer := false
		for _, l := range layerfilter {
			if (block.Pos.Y*16) >= l.From && (block.Pos.Y*16) <= l.To {
				inLayer = true
				break
			}
		}

		if !inLayer {
			continue
		}

		fields := logrus.Fields{
			"x": block.Pos.X,
			"y": block.Pos.Y,
			"z": block.Pos.Z,
		}
		logrus.WithFields(fields).Debug("mapblock")

		key := getKey(block.Pos)

		mapblock, err := mapblockparser.Parse(block.Data, block.Mtime, block.Pos)
		if err != nil {
			return nil, err
		}

		a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, mapblock)

		a.c.Set(key, mapblock, cache.DefaultExpiration)
		mblist = append(mblist, mapblock)

	}

	result.LastPos = newlastpos
	result.List = mblist

	return &result, nil
}

func (a *MapBlockAccessor) FindMapBlocksByPos(lastpos *coords.MapBlockCoords, limit int, layerfilter []layer.Layer) (*FindMapBlocksResult, error) {

	fields := logrus.Fields{
		"x":     lastpos.X,
		"y":     lastpos.Y,
		"z":     lastpos.Z,
		"limit": limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByPos")

	blocks, err := a.accessor.FindLegacyBlocksByPos(lastpos, limit)

	if err != nil {
		return nil, err
	}

	result := FindMapBlocksResult{}

	mblist := make([]*mapblockparser.MapBlock, 0)
	var newlastpos *coords.MapBlockCoords
	result.HasMore = len(blocks) == limit
	result.UnfilteredCount = len(blocks)

	for _, block := range blocks {
		newlastpos = block.Pos
		if result.LastMtime < block.Mtime {
			result.LastMtime = block.Mtime
		}

		inLayer := false
		for _, l := range layerfilter {
			if (block.Pos.Y*16) >= l.From && (block.Pos.Y*16) <= l.To {
				inLayer = true
				break
			}
		}

		if !inLayer {
			continue
		}

		fields := logrus.Fields{
			"x": block.Pos.X,
			"y": block.Pos.Y,
			"z": block.Pos.Z,
		}
		logrus.WithFields(fields).Trace("mapblock")

		key := getKey(block.Pos)

		mapblock, err := mapblockparser.Parse(block.Data, block.Mtime, block.Pos)
		if err != nil {
			return nil, err
		}

		a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, mapblock)

		a.c.Set(key, mapblock, cache.DefaultExpiration)
		mblist = append(mblist, mapblock)

	}

	result.LastPos = newlastpos
	result.List = mblist

	return &result, nil
}

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
