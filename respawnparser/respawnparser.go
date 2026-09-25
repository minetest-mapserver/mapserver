package respawnparser
// ref https://github.com/minetest-go/areasparser/blob/master/parser.go
// duck.ai
// docker go cache help: https://oneuptime.com/blog/post/2026-02-08-how-to-speed-up-docker-build-for-go-projects/view

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type File map[string]RespawnPlace

type RespawnPlace struct {
	FullName string     `json:"full_name,omitempty"`
	Pos      GenericPos `json:"pos"`
	Look     *Direction `json:"look,omitempty"`
	Color    string     `json:"color,omitempty"`
	Icon     string     `json:"icon,omitempty"`
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

	return nil
}

func Parse(data []byte) (File, error) {
    decoder := json.NewDecoder(bytes.NewReader(data))
    var respawnplaces File
    if err := decoder.Decode(&respawnplaces); err != nil {
        return nil, fmt.Errorf("parse respawn data: %w", err)
    }
    return respawnplaces, nil
}

// ParseFile reads a file and passes its contents to Parse.
func ParseFile(filename string) (File, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	respawnplaces, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}

	return respawnplaces, nil
}
