package colormapping

import (
	"testing"
)

func TestNewMapping(t *testing.T) {
  m := CreateColorMapping()
  err := m.LoadVFSColors(false, "/colors.txt")
  if err != nil {
    t.Fatal(err)
  }

  c := m.GetColor("scifi_nodes:blacktile2")
  if c == nil {
    panic("no color")
  }

}
