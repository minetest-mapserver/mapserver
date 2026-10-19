package mapblockaccessor

import (
	"fmt"
	"mapserver/types"
	"mapserver/util"
	"sync"

	"github.com/minetest-go/mapparser"
	"github.com/prometheus/client_golang/prometheus"
)

var lock = &sync.RWMutex{}

func (a *MapBlockAccessor) GetMapBlock(pos *types.MapBlockCoords) (*mapparser.MapBlock, error) {
	key := getKey(pos)

	cachedblock, found := a.blockcache.Get(key)
	if found {
		getCacheHitCount.Inc()
		return cachedblock, nil
	}
	getCacheMissCount.Inc()

	timer := prometheus.NewTimer(dbGetDuration)
	defer timer.ObserveDuration()

	block, err := a.accessor.GetBlock(pos)
	if err != nil {
		return nil, fmt.Errorf("mapblock accessor error: %v", err)
	}

	if block == nil {
		// no mapblock here
		a.blockcache.Add(key, nil)
		return nil, nil
	}

	mapblock, err := mapparser.Parse(block.Data)
	if err != nil {
		return nil, err
	}

	a.Eventbus.Emit(util.MAPBLOCK_RENDERED, types.NewParsedMapblock(mapblock, pos))
	a.blockcache.Add(key, mapblock)

	return mapblock, nil
}
