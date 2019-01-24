package app

import (
	"encoding/json"
	"io/ioutil"
	"mapserver/coords"
	"mapserver/layer"
	"os"
	"runtime"
	"sync"
)

type Config struct {
	Port                int              `json:"port"`
	EnableRendering     bool             `json:"enablerendering"`
	Webdev              bool             `json:"webdev"`
	WebApi              *WebApiConfig    `json:"webapi"`
	RenderState         *RenderStateType `json:"renderstate"`
	Layers              []layer.Layer    `json:"layers"`
	RenderingFetchLimit int              `json:"renderingfetchlimit"`
	RenderingJobs       int              `json:"renderingjobs"`
	RenderingQueue      int              `json:"renderingqueue"`
}

type WebApiConfig struct {
	EnableMapblock bool `json:"enablemapblock"`
}

type RenderStateType struct {
	//Initial rendering flag (true=still active)
	InitialRun      bool `json:"initialrun"`
	LegacyProcessed int  `json:"legacyprocessed"`

	//Last initial rendering coords
	LastX int `json:"lastx"`
	LastY int `json:"lasty"`
	LastZ int `json:"lastz"`

	//Last mtime of incremental rendering
	LastMtime int64 `json:"lastmtime"`
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
	}

	rstate := RenderStateType{
		InitialRun: true,
		LastX:      coords.MinCoord - 1,
		LastY:      coords.MinCoord - 1,
		LastZ:      coords.MinCoord - 1,
		LastMtime:  0,
	}

	layers := []layer.Layer{
		layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	cfg := Config{
		Port:                8080,
		EnableRendering:     true,
		Webdev:              false,
		WebApi:              &webapi,
		RenderState:         &rstate,
		Layers:              layers,
		RenderingFetchLimit: 1000,
		RenderingJobs:       runtime.NumCPU(),
		RenderingQueue:      100,
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
