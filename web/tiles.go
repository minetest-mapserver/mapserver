package web

import (
  "mapserver/app"
  "mapserver/coords"
  "net/http"
  "strings"
  "strconv"
)

type Tiles struct {
  ctx *app.App
}

func (t *Tiles) ServeHTTP(resp http.ResponseWriter, req *http.Request){
  str := strings.TrimPrefix(req.URL.Path, "/tile/")
  parts := strings.Split(str, "/")
  layerid, _ := strconv.Atoi(parts[0])
  x, _ := strconv.Atoi(parts[1])
  y, _ := strconv.Atoi(parts[2])
  zoom, _ := strconv.Atoi(parts[3])

  c := coords.NewTileCoords(x,y,zoom,layerid)
  data, err := t.ctx.Tilerenderer.Render(c)

  if err != nil {
    resp.WriteHeader(500)
    resp.Write([]byte(err.Error()))

  } else {
    resp.Header().Add("content-type", "image/png")
    resp.Write(data)
  }
}
