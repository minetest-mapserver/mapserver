package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/types"
	"net/http"
)

// Public facing config
type PublicConfig struct {
	Version         string               `json:"version"`
	Layers          []*types.Layer       `json:"layers"`
	MapObjects      *app.MapObjectConfig `json:"mapobjects"`
	DefaultOverlays []string             `json:"defaultoverlays"`
	PageName        string               `json:"pagename"`
	EnableSearch    bool                 `json:"enablesearch"`
}

func (api *Api) GetConfig(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")

	webcfg := PublicConfig{}
	webcfg.Layers = api.Context.Config.Layers
	webcfg.MapObjects = api.Context.Config.MapObjects
	webcfg.Version = app.Version
	webcfg.DefaultOverlays = api.Context.Config.DefaultOverlays
	webcfg.PageName = api.Context.Config.PageName
	webcfg.EnableSearch = api.Context.Config.EnableSearch

	json.NewEncoder(resp).Encode(webcfg)
}
