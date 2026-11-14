package mapblockaccessor

import (
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

func (a *MapBlockAccessor) Update(pos *types.MapBlockCoords, mb *mapparser.MapBlock) {
	key := getKey(pos)
	cacheBlockCount.Inc()
	a.blockcache.Add(key, mb)
}
