package colormapping

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"image/color"
	"mapserver/vfs"
	"strconv"
	"strings"
)

type ColorMapping struct {
	colors map[string]*color.RGBA
}

func (m *ColorMapping) GetColor(name string) *color.RGBA {
	return m.colors[name]
}

func (m *ColorMapping) LoadBytes(buffer []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(buffer))
	for scanner.Scan() {
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
			return errors.New("invalid line")
		}

		if len(parts) >= 4 {
			r, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				return err
			}

			g, err := strconv.ParseInt(parts[2], 10, 32)
			if err != nil {
				return err
			}

			b, err := strconv.ParseInt(parts[3], 10, 32)
			if err != nil {
				return err
			}

			c := color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}
			m.colors[parts[0]] = &c
		}

		if len(parts) >= 5 {
			//with alpha
		}

	}

	return nil
}

//TODO: colors from fs

func (m *ColorMapping) LoadVFSColors(useLocal bool, filename string) error {
	buffer, err := vfs.FSByte(useLocal, "/colors.txt")
	if err != nil {
		return err
	}

	log.WithFields(logrus.Fields{"size": len(buffer),
		"filename": filename,
		"useLocal": useLocal}).Info("Loading local colors file")

	return m.LoadBytes(buffer)
}

func NewColorMapping() *ColorMapping {
	return &ColorMapping{colors: make(map[string]*color.RGBA)}
}
