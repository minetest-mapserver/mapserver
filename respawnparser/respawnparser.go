package respawnparser
// ref https://github.com/minetest-go/areasparser/blob/master/parser.go
// duck.ai
// docker go cache help: https://oneuptime.com/blog/post/2026-02-08-how-to-speed-up-docker-build-for-go-projects/view

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type File map[string]RespawnPlace

type RespawnPlace struct {
	FullName string     `json:"full_name,omitempty"`
	Pos      GenericPos `json:"pos"`
	Look     *Direction `json:"look,omitempty"`
}

type GenericPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Direction struct {
	H float64 `json:"h"`
	V float64 `json:"v"`
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
    pos.X = getInt(m["x"])
    pos.Y = getInt(m["y"])
    pos.Z = getInt(m["z"])
	// do not cast to int like https://github.com/minetest-go/areasparser/blob/master/parser.go
	return nil
}

func ParseFile(filename string) (File, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var respawnplaces File
	if err := json.Unmarshal(data, &respawnplaces); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}

	return respawnplaces, nil
}
