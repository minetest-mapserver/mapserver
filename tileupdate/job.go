package tileupdate

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"time"
)

func Job(ctx *app.App) {
	rstate := ctx.Config.RenderState

	fields := logrus.Fields{
		"lastmtime": rstate.LastMtime,
	}
	logrus.WithFields(fields).Info("Starting incremental update")

	for true {
		mblist, err := ctx.BlockAccessor.FindLatestMapBlocks(rstate.LastMtime, 1000)

		if err != nil {
			panic(err)
		}

		for _, mb := range mblist {
			if mb.Mtime > rstate.LastMtime {
				rstate.LastMtime = mb.Mtime + 1
			}
		}

		ctx.Config.Save()

		fields = logrus.Fields{
			"count":     len(mblist),
			"lastmtime": rstate.LastMtime,
		}
		logrus.WithFields(fields).Info("incremental update")

		time.Sleep(5 * time.Second)
	}
}
