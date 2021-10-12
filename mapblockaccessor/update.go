package mapblockaccessor

import (
	"mapserver/coords"

	"github.com/minetest-go/mapparser"
	cache "github.com/patrickmn/go-cache"
)

func (a *MapBlockAccessor) Update(pos *coords.MapBlockCoords, mb *mapparser.MapBlock) {
	key := getKey(pos)
	cacheBlockCount.Inc()
	a.blockcache.Set(key, mb, cache.DefaultExpiration)
}
