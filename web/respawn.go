package web

import (
	"encoding/json"
	"net/http"

	"os"
	"sync"
	"time"

	"mapserver/app"
	"mapserver/respawnparser"
)

type RespawnPlacesHandler struct {
	ctx      *app.App
	cache    map[string]respawnparser.RespawnPlace
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

	now := time.Now().Unix()
	diff := now - h.lasttime

	if diff > 5 {
		mutex_respawn.Lock()

		h.lasttime = now
		places, err := respawnparser.ParseFile(RESPAWN_PLACES_FILENAME)

		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}

		h.cache = places

		mutex_respawn.Unlock()
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(h.cache)

}
