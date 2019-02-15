package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapobjectdb"
	"net/http"
	"strconv"
	"strings"
	"github.com/prometheus/client_golang/prometheus"
)

type MapObjects struct {
	ctx *app.App
}

func (t *MapObjects) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/mapobjects/")
	// x1/y1/z1/x2/y2/z2/type
	parts := strings.Split(str, "/")
	if len(parts) != 7 {
		resp.WriteHeader(500)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	timer := prometheus.NewTimer(mapobjectServeDuration)
	defer timer.ObserveDuration()

	x1, _ := strconv.Atoi(parts[0])
	y1, _ := strconv.Atoi(parts[1])
	z1, _ := strconv.Atoi(parts[2])
	x2, _ := strconv.Atoi(parts[3])
	y2, _ := strconv.Atoi(parts[4])
	z2, _ := strconv.Atoi(parts[5])
	typeStr := parts[6]

	q := mapobjectdb.SearchQuery{
		Pos1: coords.NewMapBlockCoords(x1, y1, z1),
		Pos2: coords.NewMapBlockCoords(x2, y2, z2),
		Type: typeStr,
	}
	objects, err := t.ctx.Objectdb.GetMapData(q)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {
		resp.Header().Add("content-type", "application/json")
		json.NewEncoder(resp).Encode(objects)

	}
}
