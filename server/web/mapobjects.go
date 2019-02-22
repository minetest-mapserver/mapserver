package web

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"mapserver/app"
	"mapserver/mapobjectdb"
	"net/http"
)

type MapObjects struct {
	ctx *app.App
}

func (t *MapObjects) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	timer := prometheus.NewTimer(mapobjectServeDuration)
	defer timer.ObserveDuration()

	decoder := json.NewDecoder(req.Body)
	var q mapobjectdb.SearchQuery

	err := decoder.Decode(&q)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	objects, err := t.ctx.Objectdb.GetMapData(q)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(objects)
}
