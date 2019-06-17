package web

import (
	"mapserver/app"
	"mapserver/vfs"
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

	tiles := &Tiles{ctx: ctx}
	tiles.Init()
	mux.Handle("/api/tile/", tiles)
	mux.Handle("/api/config", &ConfigHandler{ctx: ctx})
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI

		if len(uri) >= 3 {
			suffix := uri[len(uri)-3:]

			switch suffix {
			case "css":
				w.Header().Set("Content-Type", "text/css")
			case ".js":
				w.Header().Set("Content-Type", "application/javascript")
			case "png":
				w.Header().Set("Content-Type", "image/png")
			}
		}
		mux.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(ctx.Config.Port), nil)
	if err != nil {
		panic(err)
	}
}
