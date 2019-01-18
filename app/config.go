package app

import (
  "encoding/json"
  "io/ioutil"
  "os"
)

type Config struct {
  Port int        `json:"port"`
}

func ParseConfig(filename string) (*Config, error) {
  cfg := Config{
    Port: 80,
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
