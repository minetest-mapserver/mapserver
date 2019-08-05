package colormapping

import (
	"bytes"
	"image/color"
	"image/png"
)

type Palette struct {
	colors map[int]*color.RGBA
}

func NewPalette(imagefile []byte) (*Palette, error) {
	palette := &Palette{
		colors: make(map[int]*color.RGBA),
	}

	reader := bytes.NewReader(imagefile)
	img, err := png.Decode(reader)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	index := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()

			//fmt.Println("x ", x, " y ", y, " Index: ", index, " Color ", c)
			palette.colors[index] = &color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

			index++
		}
	}

	return palette, nil
}

func (m *Palette) GetColor(param2 int) *color.RGBA {
	return m.colors[param2]
}
