package web

import (
	"encoding/json"
	"mapserver/types"
	"net/http"
	"strconv"
	"strings"
)

type ViewBlock struct {
	BlockMapping map[int]string `json:"blockmapping"`
	ContentId    []int          `json:"contentid"`
}

func (api *Api) GetBlockData(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/viewblock/")
	parts := strings.Split(str, "/")
	if len(parts) != 3 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	x, errx := strconv.Atoi(parts[0])
	y, erry := strconv.Atoi(parts[1])
	z, errz := strconv.Atoi(parts[2])
	if errx != nil || erry != nil || errz != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid coordinates"))
		return
	}

	c := types.NewMapBlockCoords(x, y, z)
	mb, err := api.Context.MapBlockAccessor.GetMapBlock(c)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {

		var vb *ViewBlock
		if mb != nil {
			vb = &ViewBlock{}
			vb.BlockMapping = mb.BlockMapping
			vb.ContentId = mb.Mapdata.ContentId
		}

		resp.Header().Add("content-type", "application/json")
		json.NewEncoder(resp).Encode(vb)

	}
}
