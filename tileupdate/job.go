package tileupdate

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"time"
)

func Job(ctx *app.App) {
	//TODO remember last time run
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

		for _, mb := range mblist {
			if mb.Mtime > t {
				t = mb.Mtime + 1
			}
		}

		fields = logrus.Fields{
			"count": len(mblist),
			"time":  t,
		}
		logrus.WithFields(fields).Info("incremental update")

		time.Sleep(5 * time.Second)
	}
}
