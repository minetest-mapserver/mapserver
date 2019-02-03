package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"

	"github.com/sirupsen/logrus"
)

type PlayerPos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Player struct {
	Pos    PlayerPos `json:"pos"`
	Name   string    `json:"name"`
	HP     int       `json:"hp"`
	Breath int       `json:"breath"`
	//TODO: stamina, skin, etc
}

type MinetestInfo struct {
	MaxLag  float64  `json:"max_lag"`
	Players []Player `json:"players"`
	Time    float64  `json:"time"`
	Uptime  float64  `json:"uptime"`
}

type Minetest struct {
	ctx *app.App
}

func (t *Minetest) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	data := &MinetestInfo{}

	err := json.NewDecoder(req.Body).Decode(data)
	logrus.Info(data)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		logrus.Warn(err)
		return
	}

	json.NewEncoder(resp).Encode("stub")
}
