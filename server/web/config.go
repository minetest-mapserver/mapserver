package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/layer"
	"net/http"
)

//Public facing config
type PublicConfig struct {
	Layers     []*layer.Layer       `json:"layers"`
	MapObjects *app.MapObjectConfig `json:"mapobjects"`
}

type ConfigHandler struct {
	ctx *app.App
}

func (h *ConfigHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")

	webcfg := PublicConfig{}
	webcfg.Layers = h.ctx.Config.Layers
	webcfg.MapObjects = h.ctx.Config.MapObjects

	json.NewEncoder(resp).Encode(webcfg)
}
