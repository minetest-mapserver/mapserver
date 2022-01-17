package web

import (
	"encoding/json"
	"mapserver/coords"
	"mapserver/mapobjectdb"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func (api *Api) QueryMapobjects(resp http.ResponseWriter, req *http.Request) {

	timer := prometheus.NewTimer(mapobjectServeDuration)
	defer timer.ObserveDuration()

	decoder := json.NewDecoder(req.Body)
	q := mapobjectdb.SearchQuery{}

	err := decoder.Decode(&q)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	// apply defaults
	limit := 1000
	if q.Limit != nil {
		q.Limit = &limit
	}

	if q.Pos1 == nil {
		q.Pos1 = &coords.MapBlockCoords{X: -2048, Y: -2048, Z: -2048}
	}

	if q.Pos2 == nil {
		q.Pos2 = &coords.MapBlockCoords{X: 2048, Y: 2048, Z: 2048}
	}

	objects, err := api.Context.Objectdb.GetMapData(&q)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(objects)
}
