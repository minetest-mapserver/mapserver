package web

import (
	"embed"
	"mapserver/app"
	"mapserver/public"
	"net/http"
	"os"
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

	// static files
	mux.Handle("/", http.FileServer(getFileSystem(ctx.Config.Webdev, public.Files)))

	tiles := &Tiles{ctx: ctx}
	tiles.Init()
	mux.Handle("/api/tile/", tiles)
	mux.Handle("/api/config", &ConfigHandler{ctx: ctx})
	mux.Handle("/api/stats", &StatsHandler{ctx: ctx})
	mux.Handle("/api/media/", &MediaHandler{ctx: ctx})
	mux.Handle("/api/minetest", &Minetest{ctx: ctx})
	mux.Handle("/api/mapobjects/", &MapObjects{ctx: ctx})
	mux.Handle("/api/colormapping", &ColorMappingHandler{ctx: ctx})
	mux.Handle("/api/viewblock/", &ViewMapblockHandler{ctx: ctx})

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

	http.HandleFunc("/", mux.ServeHTTP)

	err := http.ListenAndServe(":"+strconv.Itoa(ctx.Config.Port), nil)
	if err != nil {
		panic(err)
	}
}

func getFileSystem(useLocalfs bool, content embed.FS) http.FileSystem {
	if useLocalfs {
		log.Print("using live mode")
		return http.FS(os.DirFS("public"))
	}

	log.Print("using embed mode")
	return http.FS(content)
}
