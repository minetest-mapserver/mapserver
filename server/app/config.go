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
	Port                int           `json:"port"`
	EnablePrometheus    bool          `json:"enableprometheus"`
	EnableRendering     bool          `json:"enablerendering"`
	Webdev              bool          `json:"webdev"`
	WebApi              *WebApiConfig `json:"webapi"`
	Layers              []layer.Layer `json:"layers"`
	RenderingFetchLimit int           `json:"renderingfetchlimit"`
	RenderingJobs       int           `json:"renderingjobs"`
	RenderingQueue      int           `json:"renderingqueue"`
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
		EnablePrometheus:    true,
		Webdev:              false,
		WebApi:              &webapi,
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
