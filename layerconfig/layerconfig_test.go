package layerconfig

import (
  "testing"
)

func TestReadJson(t *testing.T){
  layers, err := ParseFile("./testdata/layers.json")

  if err != nil {
    t.Fatal(err)
  }

  if layers == nil {
    t.Fatal("no data")
  }

  if len(layers) != 1 {
    t.Fatal("length mismatch")
  }

  if layers[0].Name != "Base" {
    t.Fatal("name mismatch")
  }

}
