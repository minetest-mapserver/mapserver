package mapblockparser

import (
	"testing"
	"io/ioutil"
	"strconv"
	mapblockparser "mapserver/mapblockparser"
)

func TestParse(t *testing.T){
	data, err := ioutil.ReadFile("testdata/0.0.0")
	if err != nil {
		t.Error(err)
	}

	mapblock, err := mapblockparser.Parse(data)

	if err != nil {
		t.Error(err)
	}

	if mapblock.Version != 28 {
		t.Error("wrong mapblock version: " + strconv.Itoa(int(mapblock.Version)))
	}

	if !mapblock.Underground {
		t.Error("Underground flag")
	}

	if len(mapblock.Mapdata) != 16384 {
		t.Error("Mapdata length wrong")
	}
}