package colormapping

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewPalette(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/unifieddyes_palette_extended.png")

	if err != nil {
		t.Fatal(err)
	}

	palette, err := NewPalette(data)

	if err != nil {
		t.Fatal(err)
	}

	color := palette.GetColor(0)

	if color == nil {
		t.Fatal("color not found!")
	}

	fmt.Println(color)

}
