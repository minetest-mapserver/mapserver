package web

import (
	"encoding/json"
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

	objects, err := api.Context.Objectdb.GetMapData(&q)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(objects)
}
