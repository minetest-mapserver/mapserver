package mapblockaccessor

import (
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

func (a *MapBlockAccessor) Update(pos *types.MapBlockCoords, mb *mtypes.MapBlock) {
	key := getKey(pos)
	cacheBlockCount.Inc()
	a.blockcache.Add(key, mb)
}
