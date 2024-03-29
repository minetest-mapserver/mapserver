package mapblockaccessor

import (
	"fmt"
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/types"

	"time"

	cache "github.com/patrickmn/go-cache"
)

type MapBlockAccessor struct {
	accessor   db.DBAccessor
	blockcache *cache.Cache
	Eventbus   *eventbus.Eventbus
	maxcount   int
}

func getKey(pos *types.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor, expiretime, purgetime time.Duration, maxcount int) *MapBlockAccessor {
	blockcache := cache.New(expiretime, purgetime)

	return &MapBlockAccessor{
		accessor:   accessor,
		blockcache: blockcache,
		Eventbus:   eventbus.New(),
		maxcount:   maxcount,
	}
}
