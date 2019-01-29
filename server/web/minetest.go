package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"
)

type Minetest struct {
	ctx *app.App
}

func (t *Minetest) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode("stub")

}
