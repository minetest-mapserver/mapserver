package layerconfig

import (
  "encoding/json"
  "io/ioutil"
)

type LayerContainer struct {
  Layers []Layer `json:"layers"`
}

type Layer struct {
  Id int        `json:"id"`
  Name string  `json:"name"`
  To int        `json:"to"`
  From int  `json:"from"`
}

var DefaultLayers []Layer

func init(){
  DefaultLayers = []Layer{
    Layer{
      Id: 0,
      Name: "Base",
      From: -16,
      To: 160,
    },
  }
}

func ParseFile(filename string) ([]Layer, error) {

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  lc := LayerContainer{}

  err = json.Unmarshal(data, &lc)
  if err != nil {
    return nil, err
  }

  return lc.Layers, nil
}
