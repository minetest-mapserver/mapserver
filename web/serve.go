package web

import (
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
		"port":          ctx.Config.Port,
		"webdev":        ctx.Config.Webdev,
		"localhostonly": ctx.Config.LocalhostOnly,
	}
	logrus.WithFields(fields).Info("Starting http server")

	api := NewApi(ctx)
	mux := http.NewServeMux()

	// static files
	if ctx.Config.Webdev {
		logrus.Print("using live mode")
		fs := http.FileServer(http.FS(os.DirFS("public/dist")))
		mux.HandleFunc("GET /", fs.ServeHTTP)

	} else {
		logrus.Print("using embed mode")
		fs := http.FileServer(http.FS(public.Files))
		mux.HandleFunc("GET /", CachedServeFunc(fs.ServeHTTP))
	}

	mux.Handle("GET /api/tile/", &Tiles{ctx: ctx})
	mux.HandleFunc("GET /api/config", api.GetConfig)
	mux.HandleFunc("GET /api/stats", api.GetStats)
	mux.HandleFunc("POST /api/luanti/stats", api.PostLuantiStats)
	mux.HandleFunc("POST /api/luanti/nodedefs", api.PostLuantiNodedefs)
	mux.HandleFunc("POST /api/luanti/media", api.PostLuantiMedia)
	mux.HandleFunc("POST /api/mapobjects/", api.QueryMapobjects)
	mux.HandleFunc("GET /api/colormapping", api.GetColorMapping)
	mux.HandleFunc("GET /api/viewblock/", api.GetBlockData)

	if ctx.Config.MapObjects.Areas {
		mux.Handle("GET /api/areas", &AreasHandler{ctx: ctx})
	}

	if ctx.Config.MapObjects.RespawnPlaces {
		mux.Handle("GET /api/respawnplaces", &RespawnPlacesHandler{ctx: ctx})
	}

	if ctx.Config.EnablePrometheus {
		mux.Handle("GET /metrics", promhttp.Handler())
	}

	ws := NewWS(ctx)
	mux.Handle("GET /api/ws", ws)

	ctx.Tilerenderer.Eventbus.AddListener(ws)
	ctx.WebEventbus.AddListener(ws)

	var bind_addr string = ":" + strconv.Itoa(ctx.Config.Port)
	if ctx.Config.LocalhostOnly {
		bind_addr = "127.0.0.1:" + strconv.Itoa(ctx.Config.Port)
	}

	// main entry point
	http.HandleFunc("/", mux.ServeHTTP)
	err := http.ListenAndServe(bind_addr, nil)
	if err != nil {
		panic(err)
	}
}
