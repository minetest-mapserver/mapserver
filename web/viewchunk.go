package web

import (
	"encoding/json"
	"fmt"
	"mapserver/types"
	"net/http"
	"strconv"

	"github.com/minetest-go/mapparser"
)

type ViewBlock struct {
	//TODO: send as deflated base64 string
	Mapdata      *mapparser.MapData `json:"mapdata"`
	BlockMapping map[int]string     `json:"blockmapping"`
}

func (api *Api) ViewChunk(resp http.ResponseWriter, req *http.Request) {

	// chunk pos
	chunk_x, _ := strconv.ParseInt(req.PathValue("chunk_x"), 10, 32)
	chunk_y, _ := strconv.ParseInt(req.PathValue("chunk_y"), 10, 32)
	chunk_z, _ := strconv.ParseInt(req.PathValue("chunk_z"), 10, 32)

	blocks := []*ViewBlock{}

	x1, y1, z1, x2, y2, z2 := types.GetMapblockBoundsFromChunk(int(chunk_x), int(chunk_y), int(chunk_z))
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				coords := types.NewMapBlockCoords(x, y, z)
				mb, err := api.Context.MapBlockAccessor.GetMapBlock(coords)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
					resp.Write([]byte(err.Error()))
					return
				}
				if mb != nil {
					blocks = append(blocks, &ViewBlock{
						Mapdata:      mb.Mapdata,
						BlockMapping: mb.BlockMapping,
					})
				}
			}
		}
	}

	resp.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(resp).Encode(blocks)
	if err != nil {
		fmt.Printf("viewblock json encode error: %v", err)
	}
}
