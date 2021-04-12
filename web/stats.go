package web

import (
	"encoding/json"
	"net/http"
)

func (api *Api) GetStats(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(LastStats)
}
