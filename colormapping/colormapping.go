package colormapping

import (
	"bufio"
	"bytes"
	"errors"
	"image/color"
	"mapserver/vfs"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type ColorMapping struct {
	colors               map[string]*color.RGBA
	extendedpaletteblock map[string]bool
	extendedpalette      *Palette
}

func (m *ColorMapping) GetColor(name string, param2 int) *color.RGBA {
	//TODO: list of node->palette
	if m.extendedpaletteblock[name] {
		// param2 coloring
		return m.extendedpalette.GetColor(param2)
	}

	return m.colors[name]
}

func (m *ColorMapping) LoadBytes(buffer []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(buffer))
	count := 0
	line := 0

	for scanner.Scan() {
		line++

		txt := strings.Trim(scanner.Text(), " ")

		if len(txt) == 0 {
			//empty
			continue
		}

		if strings.HasPrefix(txt, "#") {
			//comment
			continue
		}

		parts := strings.Fields(txt)

		if len(parts) < 4 {
			return 0, errors.New("invalid line: #" + strconv.Itoa(line))
		}

		if len(parts) >= 4 {
			r, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				return 0, err
			}

			g, err := strconv.ParseInt(parts[2], 10, 32)
			if err != nil {
				return 0, err
			}

			b, err := strconv.ParseInt(parts[3], 10, 32)
			if err != nil {
				return 0, err
			}

			a := int64(255)

			if len(parts) >= 5 {
				//with alpha
				//a, err = strconv.ParseInt(parts[4], 10, 32)
				//if err != nil {
				//	return 0, err
				//}
			}

			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			m.colors[parts[0]] = &c
			count++
		}
	}

	return count, nil
}

func (m *ColorMapping) LoadVFSColors(useLocal bool, filename string) (int, error) {
	buffer, err := vfs.FSByte(useLocal, "/colors.txt")
	if err != nil {
		return 0, err
	}

	log.WithFields(logrus.Fields{"size": len(buffer),
		"filename": filename,
		"useLocal": useLocal}).Info("Loading colors")

	return m.LoadBytes(buffer)
}

func NewColorMapping() *ColorMapping {
	extendedpalette, err := NewPalette(vfs.FSMustByte(false, "/pics/unifieddyes_palette_extended.png"))

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(vfs.FSMustByte(false, "/extended_palette.txt")))
	extendedpaletteblock := make(map[string]bool)

	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		txt := strings.Trim(scanner.Text(), " ")

		if len(txt) == 0 {
			//empty
			continue
		}

		extendedpaletteblock[txt] = true
	}

	return &ColorMapping{
		colors:               make(map[string]*color.RGBA),
		extendedpaletteblock: extendedpaletteblock,
		extendedpalette:      extendedpalette,
	}
}
