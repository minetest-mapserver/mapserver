package travelnetparser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mapserver/luaparser"
)

type GenericPos struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Travelnet struct {
	Timestamp int64       `json:"timestamp"`
	Pos       *GenericPos `json:"pos"`
}

func ParseFile(filename string) (map[string]map[string]map[string]*Travelnet, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(content)
}

// owner -> network -> name ->Travelnet-Data
func Parse(data []byte) (map[string]map[string]map[string]*Travelnet, error) {
	p := luaparser.New()
	travelnets := make(map[string]map[string]map[string]*Travelnet)

	ownermap, err := p.ParseMap(string(data[:]))

	if err != nil {
		return nil, err
	}

	for owner, rawnetworkmap := range ownermap {
		fmt.Println(owner)

		networkmap, ok := rawnetworkmap.(map[string]interface{})
		if !ok {
			return nil, errors.New("networkmap parsing failed")
		}

		for network, rawstationmap := range networkmap {
			fmt.Println("+", network)

			stationmap, ok := rawstationmap.(map[string]interface{})
			if !ok {
				return nil, errors.New("stationmap parsing failed")
			}

			for station, rawentries := range stationmap {
				fmt.Println("++", station)

				entries, ok := rawentries.(map[string]interface{})
				if !ok {
					return nil, errors.New("entries parsing failed")
				}

				for entry := range entries {
					fmt.Println("+++", entry)
					//TODO:
				}

			}
		}
	}

	return travelnets, nil
}
