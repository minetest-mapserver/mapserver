package web

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/vfs"
	"net/http"
	"strconv"
)

func Serve(ctx *app.App) {
	fields := logrus.Fields{
		"port":   ctx.Config.Port,
		"webdev": ctx.Config.Webdev,
	}
	logrus.WithFields(fields).Info("Starting http server")

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(vfs.FS(ctx.Config.Webdev)))

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Api endpoint")
	})

	err := http.ListenAndServe(":"+strconv.Itoa(ctx.Config.Port), mux)
	if err != nil {
		panic(err)
	}
}
