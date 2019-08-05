package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/coords"
	"net/http"
	"strconv"
	"strings"
)

type MapblockHandler struct {
	ctx *app.App
}

func (h *MapblockHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/mapblock/")
	parts := strings.Split(str, "/")
	if len(parts) != 3 {
		resp.WriteHeader(500)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	c := coords.NewMapBlockCoords(x, y, z)
	mb, err := h.ctx.MapBlockAccessor.GetMapBlock(c)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {
		resp.Header().Add("content-type", "application/json")
		json.NewEncoder(resp).Encode(mb)

	}
}
