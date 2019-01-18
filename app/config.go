package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Port                    int          `json:"port"`
	EnableInitialRendering  bool         `json:"enableinitialrendering"`
	EnableIncrementalUpdate bool         `json:"enableincrementalupdate"`
	Webdev                  bool         `json:"webdev"`
	WebApi                  WebApiConfig `json:"webapi"`
}

type WebApiConfig struct {
	EnableMapblock bool `json:"enablemapblock"`
}

func ParseConfig(filename string) (*Config, error) {
	webapi := WebApiConfig{
		EnableMapblock: false,
	}

	cfg := Config{
		Port:                    8080,
		EnableInitialRendering:  true,
		EnableIncrementalUpdate: true,
		Webdev:                  false,
		WebApi:                  webapi,
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
