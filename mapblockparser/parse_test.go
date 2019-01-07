package mapblockparser

import (
	"fmt"
	"testing"
	"io/ioutil"
	"strconv"
	log "github.com/sirupsen/logrus"
)

func TestReadU16(t *testing.T){
	v := readU16([]byte{0x00, 0x00}, 0)
	if v != 0 {
		t.Error(v)
	}

	v = readU16([]byte{0x00, 0x01}, 0)
	if v != 1 {
		t.Error(v)
	}

	v = readU16([]byte{0x01, 0x00}, 0)
	if v != 256 {
		t.Error(v)
	}

}
func TestReadU32(t *testing.T){
	v := readU32([]byte{0x00, 0x00, 0x00, 0x00}, 0)
	if v != 0 {
		t.Error(v)
	}
}

func TestParse(t *testing.T){
	log.SetLevel(log.DebugLevel)

	data, err := ioutil.ReadFile("testdata/0.0.0")
	if err != nil {
		t.Error(err)
	}

	mapblock, err := Parse(data)
	fmt.Println("mapblock.Metadata", mapblock.Metadata)

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

	pairs := mapblock.Metadata.GetPairsMap(0)
	if pairs["owner"] != "pipo" {
		t.Error(pairs["owner"])
	}
}


func TestParse2(t *testing.T){
	log.SetLevel(log.DebugLevel)

	data, err := ioutil.ReadFile("testdata/0.9.0")
	if err != nil {
		t.Error(err)
	}

	_, err = Parse(data)
	//fmt.Println("mapblock.Metadata", mapblock.Metadata)

	if err != nil {
		t.Error(err)
	}
}
