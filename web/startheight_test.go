package web

import (
	"mapserver/app"
	"mapserver/db/sqlite"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"mapserver/types"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"

	mtypes "github.com/minetest-go/types"
	"github.com/stretchr/testify/assert"
)

func TestStartHeight(t *testing.T) {
	dbfile := path.Join(t.TempDir(), "map.sqlite")
	assert.NoError(t, testutils.CreateTestDatabase(dbfile))
	a, err := sqlite.New(dbfile)
	assert.NoError(t, err)
	assert.NoError(t, a.Migrate())

	ctx := &app.App{
		Config: &app.Config{
			Layers: []*types.Layer{
				{Id: 0, Name: "full", From: -8, To: 8},
				{Id: 1, Name: "air", From: 5, To: 7},
			},
		},
		MapBlockAccessor: mapblockaccessor.NewMapBlockAccessor(a, time.Second, 1000),
	}
	api := NewApi(ctx)

	get := func(url string) *httptest.ResponseRecorder {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/startheight/{layerId}/{pos_x}/{pos_z}", api.StartHeight)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest("GET", url, nil))
		return rec
	}

	// the fixture's only non-air mapblocks contain "ignore" nodes and are skipped
	assert.Equal(t, http.StatusNotFound, get("/api/startheight/0/0/0").Code)

	// layer without non-air blocks in the column
	assert.Equal(t, http.StatusNotFound, get("/api/startheight/1/0/0").Code)
	// column without any mapblocks
	assert.Equal(t, http.StatusNotFound, get("/api/startheight/0/-3000/-3000").Code)
	// unknown layer
	assert.Equal(t, http.StatusNotFound, get("/api/startheight/99/0/0").Code)
	// invalid params
	assert.Equal(t, http.StatusBadRequest, get("/api/startheight/x/0/0").Code)
	assert.Equal(t, http.StatusBadRequest, get("/api/startheight/0/x/0").Code)
	assert.Equal(t, http.StatusBadRequest, get("/api/startheight/0/0/x").Code)
}

func TestStartHeightHelpers(t *testing.T) {
	mb := &mtypes.MapBlock{ContentId: make([]int, 4096), BlockMapping: map[int]string{0: "air", 1: "default:stone"}}
	mb.ContentId[mtypes.NewPos(3, 9, 4).Index()] = 1
	assert.False(t, containsSkippedNode(mb))
	assert.Equal(t, 9, topNodeOffset(mb, 3, 4))
	assert.Equal(t, 0, topNodeOffset(mb, 0, 0))

	mb.BlockMapping[2] = "ignore"
	assert.True(t, containsSkippedNode(mb))

	delete(mb.BlockMapping, 2)
	mb.BlockMapping[3] = "vacuum:vacuum"
	assert.True(t, containsSkippedNode(mb))
}
