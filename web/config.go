package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/layer"
	"net/http"
)

//Public facing config
type PublicConfig struct {
	Version         string               `json:"version"`
	Layers          []*layer.Layer       `json:"layers"`
	MapObjects      *app.MapObjectConfig `json:"mapobjects"`
	DefaultOverlays []string             `json:"defaultoverlays"`
	EnableSearch    bool                 `json:"enablesearch"`
}

func (api *Api) GetConfig(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")

	webcfg := PublicConfig{}
	webcfg.Layers = api.Context.Config.Layers
	webcfg.MapObjects = api.Context.Config.MapObjects
	webcfg.Version = app.Version
	webcfg.DefaultOverlays = api.Context.Config.DefaultOverlays
	webcfg.EnableSearch = api.Context.Config.EnableSearch

	json.NewEncoder(resp).Encode(webcfg)
}
