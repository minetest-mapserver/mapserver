package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"

	"os"
	"sync"
	"time"

	"github.com/minetest-go/areasparser"
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

	mutex.Lock()
	now := time.Now().Unix()
	if now-h.lasttime > 5 {
		areas, err := areasparser.ParseFile(AREAS_FILENAME)
		if err != nil {
			mutex.Unlock()
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}

		h.lasttime = now
		h.cache = areas
	}
	cache := h.cache
	mutex.Unlock()

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(cache)
}
