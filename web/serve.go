package web

import (
	"mapserver/app"
	"mapserver/bundle"
	"mapserver/vfs"
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	mux.Handle("/js/bundle.js", bundle.NewJsHandler(ctx.Config.Webdev))
	mux.Handle("/css/bundle.css", bundle.NewCSSHandler(ctx.Config.Webdev))

	tiles := &Tiles{ctx: ctx}
	tiles.Init()
	mux.Handle("/api/tile/", tiles)
	mux.Handle("/api/config", &ConfigHandler{ctx: ctx})
	mux.Handle("/api/media/", &MediaHandler{ctx: ctx})
	mux.Handle("/api/minetest", &Minetest{ctx: ctx})
	mux.Handle("/api/mapobjects/", &MapObjects{ctx: ctx})

	if ctx.Config.MapObjects.Areas {
		mux.Handle("/api/areas", &AreasHandler{ctx: ctx})
	}

	if ctx.Config.EnablePrometheus {
		mux.Handle("/metrics", promhttp.Handler())
	}

	ws := NewWS(ctx)
	mux.Handle("/api/ws", ws)

	ctx.Tilerenderer.Eventbus.AddListener(ws)
	ctx.WebEventbus.AddListener(ws)

	if ctx.Config.WebApi.EnableMapblock {
		//mapblock endpoint
		mux.Handle("/api/mapblock/", &MapblockHandler{ctx: ctx})
	}

	server := &http.Server{Handler: mux}
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(ctx.Config.Port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(l)
	if err != nil {
		panic(err)
	}
}
