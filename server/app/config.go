package app

import (
	"encoding/json"
	"io/ioutil"
	"mapserver/layer"
	"os"
	"runtime"
	"sync"
)

type Config struct {
	Port                int                     `json:"port"`
	EnablePrometheus    bool                    `json:"enableprometheus"`
	EnableRendering     bool                    `json:"enablerendering"`
	EnableTransparency  bool                    `json:"enabletransparency"`
	Webdev              bool                    `json:"webdev"`
	WebApi              *WebApiConfig           `json:"webapi"`
	Layers              []*layer.Layer          `json:"layers"`
	RenderingFetchLimit int                     `json:"renderingfetchlimit"`
	RenderingJobs       int                     `json:"renderingjobs"`
	RenderingQueue      int                     `json:"renderingqueue"`
	MapObjects          *MapObjectConfig        `json:"mapobjects"`
	MapBlockAccessorCfg *MapBlockAccessorConfig `json:"mapblockaccessor"`
}

type MapBlockAccessorConfig struct {
	Expiretime string `json:"expiretime"`
	Purgetime  string `json:"purgetime"`
	MaxItems   int    `json:"maxitems"`
}

type MapObjectConfig struct {
	Areas            bool `json:"areas"`
	Bones            bool `json:"bones"`
	Protector        bool `json:"protector"`
	XPProtector      bool `json:"xpprotector"`
	PrivProtector    bool `json:"privprotector"`
	Technic          bool `json:"technic"`
	LuaController    bool `json:"luacontroller"`
	Digiterms        bool `json:"digiterms"`
	Digilines        bool `json:"digilines"`
	Travelnet        bool `json:"travelnet"`
	Mapserver        bool `json:"mapserver"`
	TileServerLegacy bool `json:"tileserverlegacy"`
	Mission          bool `json:"mission"`
	Jumpdrive        bool `json:"jumpdrive"`
	Smartshop        bool `json:"smartshop"`
	Fancyvend        bool `json:"fancyvend"`
	ATM              bool `json:"atm"`
	Train            bool `json:"train"`
}

type WebApiConfig struct {
	//mapblock debugging
	EnableMapblock bool `json:"enablemapblock"`

	//mod http bridge secret
	SecretKey string `json:"secretkey"`
}

var lock sync.Mutex

const ConfigFile = "mapserver.json"

func (cfg *Config) Save() error {
	return WriteConfig(ConfigFile, cfg)
}

func WriteConfig(filename string, cfg *Config) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	str, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return err
	}

	f.Write(str)

	return nil
}

func ParseConfig(filename string) (*Config, error) {
	webapi := WebApiConfig{
		EnableMapblock: false,
		SecretKey:      RandStringRunes(16),
	}

	layers := []*layer.Layer{
		&layer.Layer{
			Id:   0,
			Name: "Ground",
			From: -1,
			To:   10,
		},
		&layer.Layer{
			Id:   1,
			Name: "Sky",
			From: 11,
			To:   24,
		},
	}

	mapobjs := MapObjectConfig{
		Areas:            true,
		Bones:            true,
		Protector:        true,
		XPProtector:      true,
		PrivProtector:    true,
		Technic:          true,
		LuaController:    true,
		Digiterms:        true,
		Digilines:        true,
		Travelnet:        true,
		Mapserver:        true,
		TileServerLegacy: true,
		Mission:          true,
		Jumpdrive:        true,
		Smartshop:        true,
		Fancyvend:        true,
		ATM:              true,
		Train:            true,
	}

	mapblockaccessor := MapBlockAccessorConfig{
		Expiretime: "15s",
		Purgetime:  "30s",
		MaxItems:   5000,
	}

	cfg := Config{
		Port:                8080,
		EnableRendering:     true,
		EnablePrometheus:    true,
		EnableTransparency:  true,
		Webdev:              false,
		WebApi:              &webapi,
		Layers:              layers,
		RenderingFetchLimit: 10000,
		RenderingJobs:       runtime.NumCPU(),
		RenderingQueue:      100,
		MapObjects:          &mapobjs,
		MapBlockAccessorCfg: &mapblockaccessor,
	}

	info, err := os.Stat(filename)
	if info != nil && err == nil {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
