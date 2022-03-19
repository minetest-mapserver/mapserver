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

	api := NewApi(ctx)
	mux := http.NewServeMux()

	// static files
	if ctx.Config.Webdev {
		logrus.Print("using live mode")
		fs := http.FileServer(http.FS(os.DirFS("public")))
		mux.HandleFunc("/", fs.ServeHTTP)

	} else {
		logrus.Print("using embed mode")
		fs := http.FileServer(http.FS(public.Files))
		mux.HandleFunc("/", CachedServeFunc(fs.ServeHTTP))
	}

	tiles := &Tiles{ctx: ctx}
	tiles.Init()
	mux.Handle("/api/tile/", tiles)
	mux.HandleFunc("/api/config", api.GetConfig)
	mux.HandleFunc("/api/stats", api.GetStats)
	mux.HandleFunc("/api/media/", api.GetMedia)
	mux.HandleFunc("/api/minetest", api.PostMinetestData)
	mux.HandleFunc("/api/mapobjects/", api.QueryMapobjects)
	mux.HandleFunc("/api/colormapping", api.GetColorMapping)
	mux.HandleFunc("/api/viewblock/", api.GetBlockData)

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
		mux.HandleFunc("/api/mapblock/", api.GetMapBlockData)
	}

	// main entry point
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
