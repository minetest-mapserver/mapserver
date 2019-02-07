package web

import (
	"mapserver/app"
	"mapserver/vfs"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Serve(ctx *app.App) {
	fields := logrus.Fields{
		"port":   ctx.Config.Port,
		"webdev": ctx.Config.Webdev,
	}
	logrus.WithFields(fields).Info("Starting http server")

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(vfs.FS(ctx.Config.Webdev)))

	tiles := &Tiles{ctx: ctx}
	tiles.Init()
	mux.Handle("/api/tile/", tiles)
	mux.Handle("/api/config", &ConfigHandler{ctx: ctx})
	mux.Handle("/api/minetest", &Minetest{ctx: ctx})
	mux.Handle("/api/mapobjects/", &MapObjects{ctx: ctx})

	ws := NewWS(ctx)
	mux.Handle("/api/ws", ws)

	ctx.Tilerenderer.Eventbus.AddListener(ws)
	ctx.WebEventbus.AddListener(ws)

	if ctx.Config.WebApi.EnableMapblock {
		//mapblock endpoint
		mux.Handle("/api/mapblock/", &MapblockHandler{ctx: ctx})
	}

	err := http.ListenAndServe(":"+strconv.Itoa(ctx.Config.Port), mux)
	if err != nil {
		panic(err)
	}
}
