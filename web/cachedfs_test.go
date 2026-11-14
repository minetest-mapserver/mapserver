package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCachedServeFunc(t *testing.T) {
	calls := 0
	h := CachedServeFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Write([]byte("content"))
	})

	// first request: full response with cache headers
	rec := httptest.NewRecorder()
	h(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "content", rec.Body.String())
	assert.Equal(t, "max-age=60", rec.Header().Get("Cache-Control"))
	etag := rec.Header().Get("ETag")
	require.NotEmpty(t, etag)
	assert.Equal(t, 1, calls)

	// matching etag: 304 and handler not invoked
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", etag)
	rec = httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusNotModified, rec.Code)
	assert.Empty(t, rec.Body.String())
	assert.Equal(t, 1, calls)

	// etag list containing the current one
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", `"old", `+etag)
	rec = httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusNotModified, rec.Code)

	// stale etag: full response again
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", `"stale"`)
	rec = httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 2, calls)
}
