package colormapping

import (
  "mapserver/vfs"
  "bufio"
  "errors"
  "bytes"
  "strings"
  "strconv"
  "github.com/sirupsen/logrus"
)

type Color struct {
  R,G,B,A int
}

type ColorMapping struct {
  colors map[string]*Color
}

func (m *ColorMapping) GetColor(name string) *Color {
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

    c := Color{}
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

      c.R = int(r)
      c.G = int(g)
      c.B = int(b)
    }

    if len(parts) >= 5 {
      //with alpha
      a, err := strconv.ParseInt(parts[4], 10, 32)
      if err != nil {
        return err
      }

      c.A = int(a)
    }

    m.colors[parts[0]] = &c
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

func CreateColorMapping() *ColorMapping {
  return &ColorMapping{colors: make(map[string]*Color)}
}
