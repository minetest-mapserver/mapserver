package mapblockaccessor

import (
	"fmt"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/mapblockparser"

	"time"

	cache "github.com/patrickmn/go-cache"
)

type MapBlockAccessor struct {
	accessor db.DBAccessor
	c        *cache.Cache
	Eventbus *eventbus.Eventbus
}

func getKey(pos *coords.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor) *MapBlockAccessor {
	c := cache.New(500*time.Millisecond, 1000*time.Millisecond)

	return &MapBlockAccessor{
		accessor: accessor,
		c:        c,
		Eventbus: eventbus.New(),
	}
}

type FindMapBlocksResult struct {
	HasMore         bool
	LastPos         *coords.MapBlockCoords
	LastMtime       int64
	List            []*mapblockparser.MapBlock
	UnfilteredCount int
}
