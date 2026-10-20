package util

import (
	"testing"

	"github.com/minetest-go/mapparser"
	"github.com/stretchr/testify/assert"
)

func TestSerializeMapdata(t *testing.T) {
	mapdata := &mapparser.MapData{
		ContentId: make([]int, 4096),
		Param1:    make([]int, 4096),
		Param2:    make([]int, 4096),
	}
	mapdata.ContentId[0] = 200
	mapdata.ContentId[1] = 500
	mapdata.Param1[0] = 1
	mapdata.Param2[0] = 2

	str, err := SerializeMapdata(mapdata)
	assert.True(t, str != "")
	assert.Nil(t, err)
}
