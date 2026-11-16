package web

import (
	"encoding/json"
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

	// column of mapblock x=0/z=0 (nodes 0..15), top non-air block
	rec := get("/api/startheight/0/0/0")
	assert.Equal(t, http.StatusOK, rec.Code)
	res := &StartHeightResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), res))
	assert.Equal(t, 8, res.BlockY)
	assert.GreaterOrEqual(t, res.Y, 8*16)
	assert.LessOrEqual(t, res.Y, 8*16+15)

	// other nodes in the same mapblock column give the same result
	rec = get("/api/startheight/0/15/15")
	assert.Equal(t, http.StatusOK, rec.Code)
	res2 := &StartHeightResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), res2))
	assert.Equal(t, res.BlockY, res2.BlockY)

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
