package mapblockaccessor

import (
	"fmt"
	"mapserver/db"
	"mapserver/types"
	"mapserver/util"

	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/minetest-go/mapparser"
)

type MapBlockAccessor struct {
	accessor   db.DBAccessor
	blockcache *expirable.LRU[string, *mapparser.MapBlock]
	Eventbus   *util.Eventbus
	maxcount   int
}

func getKey(pos *types.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor, expiretime time.Duration, maxcount int) *MapBlockAccessor {
	return &MapBlockAccessor{
		accessor:   accessor,
		blockcache: expirable.NewLRU[string, *mapparser.MapBlock](maxcount, nil, expiretime),
		Eventbus:   util.NewEventbus(),
		maxcount:   maxcount,
	}
}
