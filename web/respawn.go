package web

import (
	"encoding/json"
	"net/http"

	"os"
	"sync"
	"time"

	"mapserver/app"
	"mapserver/util"
)

type RespawnPlacesHandler struct {
	ctx      *app.App
	cache    map[string]util.RespawnPlace
	lasttime int64
}

var mutex_respawn = &sync.Mutex{}

const RESPAWN_PLACES_FILENAME = "places.respawn.db"

func (h *RespawnPlacesHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	info, err := os.Stat(RESPAWN_PLACES_FILENAME)
	if info == nil || err != nil {
		// no places file
		resp.Header().Add("content-type", "application/json")
		resp.Write([]byte("[]"))
		return
	}

	mutex_respawn.Lock()
	now := time.Now().Unix()
	if now-h.lasttime > 5 {
		places, err := util.ParseRespawnFile(RESPAWN_PLACES_FILENAME)
		if err != nil {
			mutex_respawn.Unlock()
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}

		h.lasttime = now
		h.cache = places
	}
	cache := h.cache
	mutex_respawn.Unlock()

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(cache)
}
