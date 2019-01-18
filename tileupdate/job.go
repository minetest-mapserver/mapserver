package tileupdate

import (
  "mapserver/app"
  "github.com/sirupsen/logrus"
  "time"
)

func Job(ctx *app.App){
  t := time.Now().Unix()

  fields := logrus.Fields{
    "time": t,
  }
  logrus.WithFields(fields).Info("Starting incremental update")

  for true {
    mblist, err := ctx.BlockAccessor.FindLatestMapBlocks(t, 1000)

    if err != nil {
      panic(err)
    }

    fields = logrus.Fields{
      "count": len(mblist),
      "time": t,
    }
  	logrus.WithFields(fields).Info("incremental update")

    for _, mb := range(mblist) {
      if mb.Mtime > t {
        t = mb.Mtime+1
      }
    }

    time.Sleep(5 * time.Second)
  }
}
