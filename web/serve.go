package web

import (
  "github.com/sirupsen/logrus"
  "mapserver/app"
  "net/http"
  "strconv"
  "fmt"
)

func Serve(ctx *app.App){
  fields := logrus.Fields{
    "port":           ctx.Config.Port,
  }
  logrus.WithFields(fields).Info("Starting http server")

  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
  })

  err := http.ListenAndServe(":" + strconv.Itoa(ctx.Config.Port), mux)
  if err != nil {
    panic(err)
  }
}
