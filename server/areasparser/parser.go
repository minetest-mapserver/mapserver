package areasparser

import (
  "io/ioutil"
  "mapserver/luaparser"
)

type GenericPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Area struct {
  Owner string `json:"owner"`
  Name string `json:"name"`
  Pos1      *GenericPos `json:"pos1"`
  Pos2      *GenericPos `json:"pos2"`
}

func ParseFile(filename string) ([]*Area, error) {
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return Parse(content)
}

func Parse(data []byte) ([]*Area, error) {
  p := luaparser.New()
  areas := make([]*Area, 0)

  list, err := p.ParseList(string(data[:]))

  if err != nil {
    return nil, err
  }

  for _, entry := range list {
    a := Area{}
    a.Name = entry["name"].(string)
    a.Owner = entry["owner"].(string)

    p1 := GenericPos{}
    pos1 := entry["pos1"].(map[string]interface{})
    p1.X = pos1["x"].(int)
    p1.Y = pos1["y"].(int)
    p1.Z = pos1["z"].(int)
    a.Pos1 = &p1

    p2 := GenericPos{}
    pos2 := entry["pos2"].(map[string]interface{})
    p2.X = pos2["x"].(int)
    p2.Y = pos2["y"].(int)
    p2.Z = pos2["z"].(int)
    a.Pos2 = &p2

    areas = append(areas, &a)
  }

  return areas, nil
}
