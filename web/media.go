package web

import (
	"mapserver/app"
	"net/http"
	"strings"
)

type MediaHandler struct {
	ctx *app.App
}

func (h *MediaHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/media/")
	parts := strings.Split(str, "/")
	if len(parts) != 1 {
		resp.WriteHeader(500)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	filename := parts[0]

	resp.WriteHeader(500)
	resp.Write([]byte(filename))
}
