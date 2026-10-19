package web

import (
	"net/http"
)

type ViewBlock struct {
	BlockMapping map[int]string `json:"blockmapping"`
	ContentId    []int          `json:"contentid"`
}

func (api *Api) ViewChunk(resp http.ResponseWriter, req *http.Request) {

	// chunk pos
	/*
		x, _ := strconv.ParseInt(req.PathValue("x"), 10, 32)
		y, _ := strconv.ParseInt(req.PathValue("y"), 10, 32)
		z, _ := strconv.ParseInt(req.PathValue("z"), 10, 32)

		x1,y1,z1,x2,y2,z2 := types.GetMapblockBoundsFromChunk(int(x), int(y), int(z))
	*/
}
