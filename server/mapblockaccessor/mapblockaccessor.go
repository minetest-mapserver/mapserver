package mapblockaccessor

import (
	"fmt"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/eventbus"

	"time"

	cache "github.com/patrickmn/go-cache"
)

type MapBlockAccessor struct {
	accessor db.DBAccessor
	c        *cache.Cache
	Eventbus *eventbus.Eventbus
	maxcount int
}

func getKey(pos *coords.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor, expiretime, purgetime time.Duration, maxcount int) *MapBlockAccessor {
	c := cache.New(expiretime, purgetime)

	return &MapBlockAccessor{
		accessor: accessor,
		c:        c,
		Eventbus: eventbus.New(),
		maxcount: maxcount,
	}
}
