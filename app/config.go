package app

import (
	"encoding/json"
	"io/ioutil"
	"mapserver/layer"
	"os"
	"runtime"
	"sync"
)

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
		Areas:              true,
		Bones:              true,
		Protector:          true,
		XPProtector:        true,
		PrivProtector:      true,
		TechnicQuarry:      true,
		TechnicSwitch:      true,
		TechnicAnchor:      true,
		TechnicReactor:     true,
		LuaController:      true,
		Digiterms:          true,
		Digilines:          true,
		Travelnet:          true,
		MapserverPlayer:    true,
		MapserverPOI:       true,
		MapserverLabel:     true,
		MapserverTrainline: true,
		MapserverBorder:    true,
		TileServerLegacy:   true,
		Mission:            true,
		Jumpdrive:          true,
		Smartshop:          true,
		Fancyvend:          true,
		ATM:                true,
		Train:              true,
		TrainSignal:        true,
		Minecart:           false,
		Locator:            false,
	}

	mapblockaccessor := MapBlockAccessorConfig{
		Expiretime: "15s",
		Purgetime:  "30s",
		MaxItems:   500,
	}

	defaultoverlays := []string{
		"mapserver_poi",
		"mapserver_label",
		"mapserver_player",
	}

	cfg := Config{
		ConfigVersion:             1,
		Port:                      8080,
		EnableRendering:           true,
		EnablePrometheus:          true,
		EnableSearch:              true,
		EnableInitialRendering:    true,
		EnableTransparency:        false,
		EnableMediaRepository:     false,
		Webdev:                    false,
		WebApi:                    &webapi,
		Layers:                    layers,
		RenderingFetchLimit:       10000,
		RenderingJobs:             runtime.NumCPU(),
		RenderingQueue:            100,
		IncrementalRenderingTimer: "5s",
		MapObjects:                &mapobjs,
		MapBlockAccessorCfg:       &mapblockaccessor,
		DefaultOverlays:           defaultoverlays,
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
