package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/layer"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type FindMapBlocksByMtimeResult struct {
	HasMore         bool
	LastPos         *coords.MapBlockCoords
	LastMtime       int64
	List            []*mapblockparser.MapBlock
	UnfilteredCount int
}

func (a *MapBlockAccessor) FindMapBlocksByMtime(lastmtime int64, limit int, layerfilter []*layer.Layer) (*FindMapBlocksByMtimeResult, error) {

	fields := logrus.Fields{
		"lastmtime": lastmtime,
		"limit":     limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByMtime")

	blocks, err := a.accessor.FindBlocksByMtime(lastmtime, limit)

	if err != nil {
		return nil, err
	}

	result := FindMapBlocksByMtimeResult{}

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
