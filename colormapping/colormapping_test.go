package colormapping

import (
	"testing"
)

func TestNewMapping(t *testing.T) {
	m := NewColorMapping()
	_, err := m.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		t.Fatal(err)
	}

	c := m.GetColor("scifi_nodes:blacktile2")
	if c == nil {
		panic("no color")
	}

	c = m.GetColor("default:river_water_flowing")
	if c == nil {
		panic("no color")
	}

	//if c.A != 128 {
	//	panic("wrong alpha")
	//}

}
