package web

import (
	"mapserver/app"
	"mapserver/public"
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
	fallback, hasfallback := req.URL.Query()["fallback"]

	content := h.ctx.MediaRepo[filename]

	if content == nil && hasfallback && len(fallback) > 0 {
		var err error
		content, err = public.Files.ReadFile("pics/" + fallback[0])
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if content != nil {
		resp.Write(content)
		resp.Header().Add("content-type", "image/png")
		return
	}

	resp.WriteHeader(404)
	resp.Write([]byte(filename))
}
