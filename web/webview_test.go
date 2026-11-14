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
