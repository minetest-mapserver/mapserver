package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func (a *MapBlockAccessor) GetMapBlock(pos *coords.MapBlockCoords) (*mapblockparser.MapBlock, error) {
	key := getKey(pos)

	if a.c.ItemCount() > a.maxcount {
		//flush cache
		fields := logrus.Fields{
			"cached items": a.c.ItemCount(),
			"maxcount":     a.maxcount,
		}
		logrus.WithFields(fields).Warn("Flushing cache")

		a.c.Flush()
	}

	cachedblock, found := a.c.Get(key)
	if found {
		getCacheHitCount.Inc()
		return cachedblock.(*mapblockparser.MapBlock), nil
	}

	getCacheMissCount.Inc()

	timer := prometheus.NewTimer(dbGetDuration)
	defer timer.ObserveDuration()

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
