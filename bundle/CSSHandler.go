package bundle

import (
	"net/http"
)

type CSSHandler struct {
	Webdev bool
	Cache  []byte
}

func NewCSSHandler(Webdev bool) *CSSHandler {
	h := &CSSHandler{Webdev: Webdev}

	if !Webdev {
		//populate cache
		manifest := getManifest(Webdev)
		h.Cache = createBundle(Webdev, manifest.Styles)
	}

	return h
}

func (h *CSSHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "text/css")

	if h.Cache == nil {
		//dev
		manifest := getManifest(h.Webdev)
		resp.Write(createBundle(h.Webdev, manifest.Styles))

	} else {
		//prod
		resp.Header().Add("Cache-Control", "public, max-age=3600")
		resp.Write(h.Cache)

	}
}
