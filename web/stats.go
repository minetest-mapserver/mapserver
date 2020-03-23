package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"
)

type StatsHandler struct {
	ctx *app.App
}

func (h *StatsHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(LastStats)
}
