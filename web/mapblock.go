package web

import (
	"encoding/json"
	"mapserver/types"
	"net/http"
	"strconv"
	"strings"
)

func (api *Api) GetMapBlockData(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/mapblock/")
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
		resp.Header().Add("content-type", "application/json")
		json.NewEncoder(resp).Encode(mb)

	}
}
