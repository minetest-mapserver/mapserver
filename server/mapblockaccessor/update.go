package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
)

func (a *MapBlockAccessor) Update(pos *coords.MapBlockCoords, mb *mapblockparser.MapBlock) {
	key := getKey(pos)
	cacheBlockCount.Inc()
	a.blockcache.Set(key, mb, cache.DefaultExpiration)
}
