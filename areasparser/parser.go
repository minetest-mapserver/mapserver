package areasparser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type GenericPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Area struct {
	Owner  string      `json:"owner"`
	Name   string      `json:"name"`
	Parent int         `json:"parent"`
	Pos1   *GenericPos `json:"pos1"`
	Pos2   *GenericPos `json:"pos2"`
}

func getInt(o interface{}) int {
	v, _ := o.(float64)
	return int(v)
}

func (pos *GenericPos) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	// float-like to int workaround
	pos.X = getInt(m["x"])
	pos.Y = getInt(m["y"])
	pos.Z = getInt(m["z"])

	return nil
}

func ParseFile(filename string) ([]*Area, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(content)
}

func Parse(data []byte) ([]*Area, error) {
	areas := make([]*Area, 0)
	json.NewDecoder(bytes.NewReader(data)).Decode(&areas)

	return areas, nil
}
