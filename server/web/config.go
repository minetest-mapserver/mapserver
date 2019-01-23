package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"
)

type ConfigHandler struct {
	ctx *app.App
}

func (h *ConfigHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(h.ctx.Config)
}
