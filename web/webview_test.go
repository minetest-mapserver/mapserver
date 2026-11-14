package web

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostViewMedia(t *testing.T) {
	old := webview_directory
	webview_directory = t.TempDir()
	defer func() { webview_directory = old }()

	assert.NoError(t, os.MkdirAll(path.Join(webview_directory, "textures"), 0777))
	assert.NoError(t, os.MkdirAll(path.Join(webview_directory, "models"), 0777))
	assert.NoError(t, os.WriteFile(path.Join(webview_directory, "textures", "a.png"), []byte("abcd"), 0666))
	assert.NoError(t, os.WriteFile(path.Join(webview_directory, "models", "b.obj"), []byte("xy"), 0666))
	assert.NoError(t, os.WriteFile(path.Join(webview_directory, "secret.txt"), []byte("s"), 0666))

	api := &Api{}
	body := `["a.png", "missing.png", "b.obj", "../secret.txt", "..", ""]`
	rec := httptest.NewRecorder()
	api.PostViewMedia(rec, httptest.NewRequest("POST", "/api/view/media", strings.NewReader(body)))
	assert.Equal(t, http.StatusOK, rec.Code)

	list := []*LuantiMedia{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &list))
	assert.Equal(t, 2, len(list))
	assert.Equal(t, "a.png", list[0].Filename)
	assert.Equal(t, "texture", list[0].Type)
	assert.Equal(t, base64.RawStdEncoding.EncodeToString([]byte("abcd")), list[0].Data)
	assert.Equal(t, "b.obj", list[1].Filename)
	assert.Equal(t, "model", list[1].Type)

	// invalid body
	rec = httptest.NewRecorder()
	api.PostViewMedia(rec, httptest.NewRequest("POST", "/api/view/media", strings.NewReader("{}")))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// too many files
	names := make([]string, maxViewMediaBatch+1)
	data, _ := json.Marshal(names)
	rec = httptest.NewRecorder()
	api.PostViewMedia(rec, httptest.NewRequest("POST", "/api/view/media", strings.NewReader(string(data))))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestPostLuantiMedia(t *testing.T) {
	old := webview_directory
	root := t.TempDir()
	webview_directory = path.Join(root, "mapserver_webview")
	defer func() { webview_directory = old }()

	api := newAuthTestApi()
	enc := func(s string) string { return base64.RawStdEncoding.EncodeToString([]byte(s)) }
	body, err := json.Marshal([]*LuantiMedia{
		{Filename: "a.png", Data: enc("abcd"), Type: "texture"},
		{Filename: "b.obj", Data: enc("xy"), Type: "model"},
		{Filename: "../../escaped.txt", Data: enc("x"), Type: "texture"},
		{Filename: "sub/nested.png", Data: enc("x"), Type: "texture"},
		{Filename: "..", Data: enc("x"), Type: "texture"},
		{Filename: "", Data: enc("x"), Type: "texture"},
	})
	assert.NoError(t, err)

	post := func(key string) *httptest.ResponseRecorder {
		req := httptest.NewRequest("POST", "/api/luanti/media", strings.NewReader(string(body)))
		if key != "" {
			req.Header.Set("Authorization", key)
		}
		rec := httptest.NewRecorder()
		api.PostLuantiMedia(rec, req)
		return rec
	}

	// no or wrong key: rejected, nothing written
	assert.Equal(t, http.StatusForbidden, post("").Code)
	assert.Equal(t, http.StatusForbidden, post("wrong").Code)
	assert.NoDirExists(t, webview_directory)

	// valid key: plain filenames are stored, everything else is skipped
	assert.Equal(t, http.StatusOK, post("s3cret").Code)

	data, err := os.ReadFile(path.Join(webview_directory, "textures", "a.png"))
	assert.NoError(t, err)
	assert.Equal(t, "abcd", string(data))
	data, err = os.ReadFile(path.Join(webview_directory, "models", "b.obj"))
	assert.NoError(t, err)
	assert.Equal(t, "xy", string(data))

	assert.NoFileExists(t, path.Join(root, "escaped.txt"))
	assert.NoFileExists(t, path.Join(webview_directory, "escaped.txt"))
	assert.NoDirExists(t, path.Join(webview_directory, "textures", "sub"))

	entries, err := os.ReadDir(path.Join(webview_directory, "textures"))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(entries))
	entries, err = os.ReadDir(root)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(entries))
}

func TestPostLuantiNodedefs(t *testing.T) {
	old := webview_nodedefs
	webview_nodedefs = path.Join(t.TempDir(), "nodedefs.json")
	defer func() { webview_nodedefs = old }()

	api := newAuthTestApi()

	// no key: rejected, nothing written
	rec := httptest.NewRecorder()
	api.PostLuantiNodedefs(rec, httptest.NewRequest("POST", "/api/luanti/nodedefs", strings.NewReader(`{"a":1}`)))
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.NoFileExists(t, webview_nodedefs)

	// valid key
	req := httptest.NewRequest("POST", "/api/luanti/nodedefs", strings.NewReader(`{"a":1}`))
	req.Header.Set("Authorization", "s3cret")
	rec = httptest.NewRecorder()
	api.PostLuantiNodedefs(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	data, err := os.ReadFile(webview_nodedefs)
	assert.NoError(t, err)
	assert.Equal(t, `{"a":1}`, string(data))
}
