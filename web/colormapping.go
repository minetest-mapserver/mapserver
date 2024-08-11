package web

import (
	"encoding/json"
	"net/http"
)

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

func (api *Api) GetColorMapping(resp http.ResponseWriter, req *http.Request) {

	cm := make(map[string]Color)

	for k, v := range api.Context.Colormapping.GetColors() {
		cm[k] = Color{R: v.R, G: v.G, B: v.B, A: v.A}
	}

	resp.Header().Add("content-type", "application/json")
	json.NewEncoder(resp).Encode(cm)
}
