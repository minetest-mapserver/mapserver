package mapblockaccessor

import (
	"fmt"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
	mtypes "github.com/minetest-go/types"
	"github.com/prometheus/client_golang/prometheus"
)

func (a *MapBlockAccessor) GetMapBlock(pos *types.MapBlockCoords) (*mtypes.MapBlock, error) {
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

	// no MAPBLOCK_RENDERED event here: this is a plain read (neighbours, 3d view, ...),
	// changed blocks are emitted by the initial and incremental rendering jobs
	a.blockcache.Add(key, mapblock)

	return mapblock, nil
}
