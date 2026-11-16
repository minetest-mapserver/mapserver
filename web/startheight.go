package web

import (
	"encoding/json"
	"mapserver/types"
	"net/http"
	"strconv"

	mtypes "github.com/luanti-golang/types"
)

type StartHeightResponse struct {
	// mapblock y position of the first non-air mapblock
	BlockY int `json:"block_y"`
	// node y position of the topmost non-air node at the requested x/z
	// inside that mapblock (the mapblock's lowest node if that column is all air)
	Y int `json:"y"`
}

// StartHeight returns the topmost non-air mapblock y-position in a layer's x/z column
func (api *Api) StartHeight(resp http.ResponseWriter, req *http.Request) {
	layerId, err := strconv.Atoi(req.PathValue("layerId"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid layerId"))
		return
	}
	pos_x, err := strconv.Atoi(req.PathValue("pos_x"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid pos_x"))
		return
	}
	pos_z, err := strconv.Atoi(req.PathValue("pos_z"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid pos_z"))
		return
	}

	layer := types.FindLayerById(api.Context.Config.Layers, layerId)
	if layer == nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("layer not found"))
		return
	}

	// node position -> mapblock position
	mbx, _, mbz := types.GetMapblockPosFromNode(pos_x, 0, pos_z)
	// node offset inside the mapblock
	offx, offz := pos_x-mbx*16, pos_z-mbz*16

	for y := layer.To; y >= layer.From; y-- {
		mb, err := api.Context.MapBlockAccessor.GetMapBlock(types.NewMapBlockCoords(mbx, y, mbz))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(err.Error()))
			return
		}
		if mb == nil || mb.AirOnly || mb.IsEmpty() || containsSkippedNode(mb) {
			continue
		}

		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(&StartHeightResponse{
			BlockY: y,
			Y:      y*16 + topNodeOffset(mb, offx, offz),
		})
		return
	}

	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte("no non-air mapblock found"))
}

// returns the y offset (0-15) of the topmost non-air node at x/z in the mapblock, 0 if there is none
func topNodeOffset(mb *mtypes.MapBlock, x, z int) int {
	for y := 15; y > 0; y-- {
		id := mb.GetNodeId(mtypes.NewPos(x, y, z))
		if name := mb.BlockMapping[id]; name != "air" {
			return y
		}
	}
	return 0
}

// returns true if the mapblock contains "ignore" nodes (not yet generated/loaded areas) or "vacuum:vacuum" nodes
func containsSkippedNode(mb *mtypes.MapBlock) bool {
	for _, name := range mb.BlockMapping {
		if name == "ignore" || name == "vacuum:vacuum" {
			return true
		}
	}
	return false
}
