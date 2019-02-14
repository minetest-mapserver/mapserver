package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/layer"
	"mapserver/settings"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type FindNextLegacyBlocksResult struct {
	HasMore         bool
	List            []*mapblockparser.MapBlock
	UnfilteredCount int
}

func (a *MapBlockAccessor) FindNextLegacyBlocks(s settings.Settings, layers []layer.Layer, limit int) (*FindNextLegacyBlocksResult, error) {

	fields := logrus.Fields{
		"x":     lastpos.X,
		"y":     lastpos.Y,
		"z":     lastpos.Z,
		"limit": limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByPos")

	nextResult, err := a.accessor.FindNextInitialBlocks(lastpos, limit)
	blocks := nextResult.List

	if err != nil {
		return nil, err
	}

	result := FindNextLegacyBlocksResult{}

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

	fields = logrus.Fields{
		"len(List)":       len(result.List),
		"unfilteredCount": result.UnfilteredCount,
		"hasMore":         result.HasMore,
		"limit":           limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByPos:Result")

	return &result, nil
}
