package mapblockaccessor

import (
	"mapserver/eventbus"
	"mapserver/layer"
	"mapserver/settings"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type FindNextLegacyBlocksResult struct {
	HasMore         bool
	List            []*types.ParsedMapblock
	UnfilteredCount int
	Progress        float64
	LastMtime       int64
}

func (a *MapBlockAccessor) FindNextLegacyBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*FindNextLegacyBlocksResult, error) {

	nextResult, err := a.accessor.FindNextInitialBlocks(s, layers, limit)

	if err != nil {
		return nil, err
	}

	blocks := nextResult.List
	result := FindNextLegacyBlocksResult{}

	mblist := make([]*types.ParsedMapblock, 0)
	result.HasMore = nextResult.HasMore
	result.UnfilteredCount = nextResult.UnfilteredCount
	result.Progress = nextResult.Progress
	result.LastMtime = nextResult.LastMtime

	for _, block := range blocks {

		fields := logrus.Fields{
			"x": block.Pos.X,
			"y": block.Pos.Y,
			"z": block.Pos.Z,
		}
		logrus.WithFields(fields).Trace("mapblock")

		key := getKey(block.Pos)

		mapblock, err := mapparser.Parse(block.Data)
		if err != nil {
			fields := logrus.Fields{
				"x":   block.Pos.X,
				"y":   block.Pos.Y,
				"z":   block.Pos.Z,
				"err": err,
			}
			logrus.WithFields(fields).Error("mapblock-pars")

			return nil, err
		}

		a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, types.NewParsedMapblock(mapblock, block.Pos))

		a.blockcache.Set(key, mapblock, cache.DefaultExpiration)
		cacheBlockCount.Inc()
		mblist = append(mblist, types.NewParsedMapblock(mapblock, block.Pos))

	}

	result.List = mblist

	fields := logrus.Fields{
		"len(List)":       len(result.List),
		"unfilteredCount": result.UnfilteredCount,
		"hasMore":         result.HasMore,
		"limit":           limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByPos:Result")

	return &result, nil
}
