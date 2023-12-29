package mapblockaccessor

import (
	"mapserver/eventbus"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
	cache "github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type FindMapBlocksByMtimeResult struct {
	HasMore         bool
	LastPos         *types.MapBlockCoords
	LastMtime       int64
	List            []*types.ParsedMapblock
	UnfilteredCount int
}

func (a *MapBlockAccessor) FindMapBlocksByMtime(lastmtime int64, limit int, layerfilter []*types.Layer) (*FindMapBlocksByMtimeResult, error) {

	fields := logrus.Fields{
		"lastmtime": lastmtime,
		"limit":     limit,
	}
	logrus.WithFields(fields).Debug("FindMapBlocksByMtime")

	timer := prometheus.NewTimer(dbGetMtimeDuration)
	blocks, err := a.accessor.FindBlocksByMtime(lastmtime, limit)
	timer.ObserveDuration()

	changedBlockCount.Add(float64(len(blocks)))

	if err != nil {
		return nil, err
	}

	result := FindMapBlocksByMtimeResult{}

	mblist := make([]*types.ParsedMapblock, 0)
	var newlastpos *types.MapBlockCoords
	result.HasMore = len(blocks) == limit
	result.UnfilteredCount = len(blocks)

	for _, block := range blocks {
		newlastpos = block.Pos
		if result.LastMtime < block.Mtime {
			result.LastMtime = block.Mtime
		}

		currentLayer := types.FindLayerByY(layerfilter, block.Pos.Y)

		if currentLayer == nil {
			continue
		}

		fields := logrus.Fields{
			"x": block.Pos.X,
			"y": block.Pos.Y,
			"z": block.Pos.Z,
		}
		logrus.WithFields(fields).Debug("mapblock")

		key := getKey(block.Pos)

		mapblock, err := mapparser.Parse(block.Data)
		if err != nil {
			fields := logrus.Fields{
				"x":   block.Pos.X,
				"y":   block.Pos.Y,
				"z":   block.Pos.Z,
				"err": err,
			}
			logrus.WithFields(fields).Error("parse error")
			continue
		}

		a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, types.NewParsedMapblock(mapblock, block.Pos))

		a.blockcache.Set(key, mapblock, cache.DefaultExpiration)
		cacheBlockCount.Inc()
		mblist = append(mblist, types.NewParsedMapblock(mapblock, block.Pos))

	}

	result.LastPos = newlastpos
	result.List = mblist

	return &result, nil
}
