package web

import (
	"encoding/json"
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

type Signal struct {
	Pos   GenericPos `json:"pos"`
	Green bool       `json:"green"`
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
	Id    float64    `json:"id"`
}

type Player struct {
	Pos             GenericPos `json:"pos"`
	Velocity        GenericPos `json:"velocity"`
	Name            string     `json:"name"`
	HP              float64    `json:"hp"`
	Breath          float64    `json:"breath"`
	Moderator       bool       `json:"moderator"`
	RTT             float64    `json:"rtt"`
	ProtocolVersion float64    `json:"protocol_version"`
	Yaw             float64    `json:"yaw"`
	Skin            string     `json:"skin"`
	//TODO: stamina, armor, etc
}

type AirUtilsPlane struct {
	Id        string     `json:"id"`
	Entity    string     `json:"entity"`
	Name      string     `json:"name"`
	Pos       GenericPos `json:"pos"`
	Owner     string     `json:"owner"`
	Driver    string     `json:"driver"`
	Passenger string     `json:"passenger"`
	Color     string     `json:"color"`
	Yaw       float64    `json:"yaw"`
}

type MinetestInfo struct {
	MaxLag         float64          `json:"max_lag"`
	Players        []*Player        `json:"players"`
	Trains         []*Train         `json:"trains"`
	Signals        []*Signal        `json:"signals"`
	Minecarts      []*Minecart      `json:"minecarts"`
	AirUtilsPlanes []*AirUtilsPlane `json:"airutils_planes"`
	Time           float64          `json:"time"`
	Uptime         float64          `json:"uptime"`
}

var LastStats *MinetestInfo

func (api *Api) PostMinetestData(resp http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Authorization") != api.Context.Config.WebApi.SecretKey {
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

	LastStats = data
	api.Context.WebEventbus.Emit("minetest-info", data)

	json.NewEncoder(resp).Encode("stub")
}
