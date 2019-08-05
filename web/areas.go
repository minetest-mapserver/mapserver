package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/areasparser"
	"net/http"

	"os"
	"sync"
	"time"
)

type AreasHandler struct {
	ctx      *app.App
	cache    []*areasparser.Area
	lasttime int64
}

var mutex = &sync.Mutex{}

const AREAS_FILENAME = "areas.dat"

func (h *AreasHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	info, err := os.Stat(AREAS_FILENAME)
	if info == nil || err != nil {
		// no areas file
		resp.Header().Add("content-type", "application/json")
		resp.Write([]byte("[]"))
		return
	}

	now := time.Now().Unix()
	diff := now - h.lasttime

	if diff > 5 {
		mutex.Lock()

		h.lasttime = now
		areas, err := areasparser.ParseFile(AREAS_FILENAME)

		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}

		h.cache = areas

		mutex.Unlock()
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(h.cache)

}
