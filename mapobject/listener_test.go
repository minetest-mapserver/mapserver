package mapobject

import (
	"mapserver/app"
	"mapserver/mapobjectdb"
	"mapserver/mapobjectdb/sqlite"
	"mapserver/types"
	"mapserver/util"
	"path/filepath"
	"testing"

	"github.com/minetest-go/mapparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordingListener struct {
	events []string
}

func (r *recordingListener) OnEvent(eventtype string, o interface{}) {
	r.events = append(r.events, eventtype)
}

type multiTestListener struct{}

func (multiTestListener) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) []*mapobjectdb.MapObject {
	list := []*mapobjectdb.MapObject{
		mapobjectdb.NewMapObject(mbpos, x, y, z, "multi"),
		mapobjectdb.NewMapObject(mbpos, x, y, z, "multi"),
	}
	for _, o := range list {
		o.Attributes["k"] = "v"
	}
	return list
}

func newTestListener(t *testing.T) (*Listener, *recordingListener) {
	db, err := sqlite.New(filepath.Join(t.TempDir(), "mapobjects.sqlite"))
	require.NoError(t, err)
	require.NoError(t, db.Migrate())

	rec := &recordingListener{}
	bus := util.NewEventbus()
	bus.AddListener(rec)

	l := &Listener{
		ctx:                  &app.App{Objectdb: db, WebEventbus: bus},
		objectlisteners:      make(map[string]MapObjectListener),
		multiobjectlisteners: make(map[string]MapMultiObjectListener),
	}
	return l, rec
}

func queryAll(t *testing.T, l *Listener, typ string) []*mapobjectdb.MapObject {
	limit := 100
	objs, err := l.ctx.Objectdb.GetMapData(&mapobjectdb.SearchQuery{
		Pos1:  types.NewMapBlockCoords(-10, -10, -10),
		Pos2:  types.NewMapBlockCoords(10, 10, 10),
		Type:  typ,
		Limit: &limit,
	})
	require.NoError(t, err)
	return objs
}

func TestListenerIgnoresOtherEvents(t *testing.T) {
	l, rec := newTestListener(t)
	l.AddMapObject("mapserver:poi", &PoiBlock{})

	// would panic on the type assertion if the event were not filtered out
	l.OnEvent(util.TILE_RENDERED, "not a mapblock")
	assert.Empty(t, rec.events)
}

func TestListenerSingleObject(t *testing.T) {
	l, rec := newTestListener(t)
	l.AddMapObject("mapserver:poi", &PoiBlock{Color: "green"})

	mb := newTestBlock("mapserver:poi", 5, 6, 7)
	setMeta(mb, 5, 6, 7, map[string]string{"name": "spot"})

	l.OnEvent(util.MAPBLOCK_RENDERED, types.NewParsedMapblock(mb, testMBPos))

	objs := queryAll(t, l, "poi")
	require.Len(t, objs, 1)
	assert.Equal(t, "spot", objs[0].Attributes["name"])
	assert.Equal(t, 16+5, objs[0].X)
	assert.Equal(t, -32+6, objs[0].Y)
	assert.Equal(t, 48+7, objs[0].Z)
	assert.Equal(t, []string{"mapobjects-cleared", "mapobject-created"}, rec.events)
}

func TestListenerMultiObject(t *testing.T) {
	l, rec := newTestListener(t)
	l.AddMapMultiObject("test:multi", multiTestListener{})

	mb := newTestBlock("test:multi", 1, 1, 1)
	l.OnEvent(util.MAPBLOCK_RENDERED, types.NewParsedMapblock(mb, testMBPos))

	assert.Len(t, queryAll(t, l, "multi"), 2)
	assert.Equal(t, []string{"mapobjects-cleared", "mapobject-created", "mapobject-created"}, rec.events)
}

func TestListenerUnregisteredNodeIgnored(t *testing.T) {
	l, rec := newTestListener(t)
	l.AddMapObject("mapserver:poi", &PoiBlock{})

	mb := newTestBlock("default:stone", 1, 1, 1)
	l.OnEvent(util.MAPBLOCK_RENDERED, types.NewParsedMapblock(mb, testMBPos))

	assert.Empty(t, queryAll(t, l, "poi"))
	assert.Equal(t, []string{"mapobjects-cleared"}, rec.events)
}

func TestListenerReplacesOldObjects(t *testing.T) {
	l, _ := newTestListener(t)
	l.AddMapObject("mapserver:poi", &PoiBlock{})

	mb := newTestBlock("mapserver:poi", 1, 1, 1)
	pmb := types.NewParsedMapblock(mb, testMBPos)
	l.OnEvent(util.MAPBLOCK_RENDERED, pmb)
	l.OnEvent(util.MAPBLOCK_RENDERED, pmb)
	assert.Len(t, queryAll(t, l, "poi"), 1)

	// node removed: re-render clears the stale object
	l.OnEvent(util.MAPBLOCK_RENDERED, types.NewParsedMapblock(newTestBlock("default:stone", 1, 1, 1), testMBPos))
	assert.Empty(t, queryAll(t, l, "poi"))
}
