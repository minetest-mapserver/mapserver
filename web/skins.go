package web

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

func (api *Api) GetSkin(resp http.ResponseWriter, req *http.Request) {
	filename := strings.TrimPrefix(req.URL.Path, "/api/skins/")
	// there should be no remaining path elements - abort if there are any - prevent escaping into FS
	if strings.Contains(filename, "/") {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	// we should only be serving PNG images
	if !strings.HasSuffix(filename, ".png") {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	filePath := api.Context.Config.Skins.SkinsPath + "/" + filename

	content, err := os.ReadFile(filePath)
	// make file not found more sensible
	if errors.Is(err, os.ErrNotExist) {
		resp.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the file content when available
	if content != nil {
		resp.Write(content)
		resp.Header().Add("content-type", "image/png")
		return
	}

	// fallback
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(filename))
}
