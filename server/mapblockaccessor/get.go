package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/mapblockparser"
	"sync"

	cache "github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var lock = &sync.RWMutex{}

func (a *MapBlockAccessor) GetMapBlock(pos *coords.MapBlockCoords) (*mapblockparser.MapBlock, error) {
	key := getKey(pos)

	//maintenance
	cacheBlocks.Set(float64(a.blockcache.ItemCount()))
	if a.blockcache.ItemCount() > a.maxcount {
		//flush cache
		fields := logrus.Fields{
			"cached items": a.blockcache.ItemCount(),
			"maxcount":     a.maxcount,
		}
		logrus.WithFields(fields).Debug("Flushing cache")

		a.blockcache.Flush()
	}

	//read section
	lock.RLock()

	cachedblock, found := a.blockcache.Get(key)
	if found {
		defer lock.RUnlock()

		getCacheHitCount.Inc()
		if cachedblock == nil {
			return nil, nil
		} else {
			return cachedblock.(*mapblockparser.MapBlock), nil
		}
	}

	//end read
	lock.RUnlock()

	timer := prometheus.NewTimer(dbGetDuration)
	defer timer.ObserveDuration()

	//write section
	lock.Lock()
	defer lock.Unlock()

	//try read
	cachedblock, found = a.blockcache.Get(key)
	if found {
		getCacheHitCount.Inc()
		if cachedblock == nil {
			return nil, nil
		} else {
			return cachedblock.(*mapblockparser.MapBlock), nil
		}
	}

	block, err := a.accessor.GetBlock(pos)
	if err != nil {
		return nil, err
	}

	if block == nil {
		//no mapblock here
		cacheBlockCount.Inc()
		a.blockcache.Set(key, nil, cache.DefaultExpiration)
		return nil, nil
	}

	getCacheMissCount.Inc()

	mapblock, err := mapblockparser.Parse(block.Data, block.Mtime, pos)
	if err != nil {
		return nil, err
	}

	a.Eventbus.Emit(eventbus.MAPBLOCK_RENDERED, mapblock)

	cacheBlockCount.Inc()
	a.blockcache.Set(key, mapblock, cache.DefaultExpiration)

	return mapblock, nil
}
