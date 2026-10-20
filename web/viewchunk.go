package web

import (
	"encoding/json"
	"fmt"
	"mapserver/types"
	"mapserver/util"
	"net/http"
	"strconv"
)

type ViewBlock struct {
	BlockPos     *types.MapBlockCoords `json:"block_pos"`
	Mapdata      string                `json:"mapdata"`
	BlockMapping map[int]string        `json:"blockmapping"`
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
				blockpos := types.NewMapBlockCoords(x, y, z)
				mb, err := api.Context.MapBlockAccessor.GetMapBlock(blockpos)
				if err != nil {
					fmt.Printf("viewblock GetMapBlock error: %v", err)
					resp.WriteHeader(http.StatusInternalServerError)
					resp.Write([]byte(err.Error()))
					return
				}
				if mb != nil {
					mapstr, err := util.SerializeMapdata(mb.Mapdata)
					if err != nil {
						fmt.Printf("viewblock SerializeMapdata error: %v", err)
						resp.WriteHeader(http.StatusInternalServerError)
						resp.Write([]byte(err.Error()))
						return
					}
					blocks = append(blocks, &ViewBlock{
						BlockPos:     blockpos,
						Mapdata:      mapstr,
						BlockMapping: mb.BlockMapping,
					})
				}
			}
		}
	}
	//TODO: caching

	resp.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(resp).Encode(blocks)
	if err != nil {
		fmt.Printf("viewblock json encode error: %v", err)
	}
}
