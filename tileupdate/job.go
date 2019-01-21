package tileupdate

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"time"
)

func Job(ctx *app.App) {
	rstate := ctx.Config.RenderState

	fields := logrus.Fields{
		"lastmtime": rstate.LastMtime,
	}
	logrus.WithFields(fields).Info("Starting incremental update")

	for true {
		mblist, err := ctx.BlockAccessor.FindLatestMapBlocks(rstate.LastMtime, ctx.Config.UpdateRenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		for _, mb := range mblist {
			if mb.Mtime > rstate.LastMtime {
				rstate.LastMtime = mb.Mtime
			}

			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			if tc == nil {
				panic("tile not in any layer")
			}

			for tc.Zoom > 1 {
				tc = tc.GetZoomedOutTile()
				ctx.Objectdb.RemoveTile(tc)
			}
		}

		//Render zoom 12-1
		for _, mb := range mblist {
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)
			for tc.Zoom > 1 {
				tc = tc.GetZoomedOutTile()

				fields = logrus.Fields{
					"X":       tc.X,
					"Y":       tc.Y,
					"Zoom":    tc.Zoom,
					"LayerId": tc.LayerId,
				}
				logrus.WithFields(fields).Debug("Dispatching tile rendering (update)")

				_, err = ctx.Tilerenderer.Render(tc)
				if err != nil {
					panic(err)
				}
			}
		}

		ctx.Config.Save()

		if len(mblist) > 0 {
			fields = logrus.Fields{
				"count":      len(mblist),
				"lastmtime":  rstate.LastMtime,
			}
			logrus.WithFields(fields).Info("incremental update")
		}

		time.Sleep(5 * time.Second)
	}
}
