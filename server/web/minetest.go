package web

import (
	"encoding/json"
	"mapserver/app"
	"net/http"

	"github.com/sirupsen/logrus"
)

type GenericPos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Wagon struct {
	Id         string  `json:"id"`
	PosInTrain float64 `json:"pos_in_train"`
	Type       string  `json:"type"`
}

type Train struct {
	Pos         GenericPos `json:"pos"`
	Id          string     `json:"id"`
	Wagons      []*Wagon   `json:"wagons"`
	OffTrack    bool       `json:"off_track"`
	Velocity    float64    `json:"velocity"`
	Line        string     `json:"line"`
	TextOutside string     `json:"text_outside"`
	TextInside  string     `json:"text_inside"`
}

type Minecart struct {
	Pos   GenericPos `json:"pos"`
	Speed GenericPos `json:"speed"`
}

type Player struct {
	Pos      GenericPos `json:"pos"`
	Velocity GenericPos `json:"velocity"`
	Name     string     `json:"name"`
	HP       float64    `json:"hp"`
	Breath   float64    `json:"breath"`
	//TODO: stamina, skin, etc
}

type MinetestInfo struct {
	MaxLag    float64     `json:"max_lag"`
	Players   []*Player   `json:"players"`
	Trains    []*Train    `json:"trains"`
	Minecarts []*Minecart `json:"minecarts"`
	Time      float64     `json:"time"`
	Uptime    float64     `json:"uptime"`
}

type Minetest struct {
	ctx *app.App
}

func (this *Minetest) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Authorization") != this.ctx.Config.WebApi.SecretKey {
		resp.WriteHeader(403)
		resp.Write([]byte("invalid key!"))
		return
	}

	resp.Header().Add("content-type", "application/json")
	data := &MinetestInfo{}

	err := json.NewDecoder(req.Body).Decode(data)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		log.WithFields(logrus.Fields{"error": err}).Error("Json unmarshal")

		return
	}

	mintestPlayers.Set(float64(len(data.Players)))
	mintestMaxLag.Set(data.MaxLag)

	this.ctx.WebEventbus.Emit("minetest-info", data)

	json.NewEncoder(resp).Encode("stub")
}
