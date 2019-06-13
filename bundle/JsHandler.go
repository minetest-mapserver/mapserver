package bundle

import (
	"net/http"
)

type JsHandler struct {
	Webdev bool
	Cache  []byte
}

func NewJsHandler(Webdev bool) *JsHandler {
	h := &JsHandler{Webdev: Webdev}

	if !Webdev {
		//populate cache
		manifest := getManifest(Webdev)
		h.Cache = createBundle(Webdev, manifest.Scripts)
	}

	return h
}

func (h *JsHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "text/javascript")

	if h.Cache == nil {
		//dev
		manifest := getManifest(h.Webdev)
		resp.Write(createBundle(h.Webdev, manifest.Scripts))

	} else {
		//prod
		resp.Header().Add("Cache-Control", "public, max-age=3600")
		resp.Write(h.Cache)

	}
}
