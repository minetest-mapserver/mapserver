package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"
)

type ColorMappingHandler struct {
	ctx *app.App
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

func (h *ColorMappingHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	cm := make(map[string]Color)

	for k, v := range h.ctx.Colormapping.GetColors() {
		cm[k] = Color{R: v.R, G: v.G, B: v.B}
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(cm)
}
