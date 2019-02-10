package mapblockaccessor

import (
	"mapserver/coords"
	"mapserver/mapblockparser"

	cache "github.com/patrickmn/go-cache"
)

func (a *MapBlockAccessor) Update(pos *coords.MapBlockCoords, mb *mapblockparser.MapBlock) {
	key := getKey(pos)
	a.c.Set(key, mb, cache.DefaultExpiration)
}
