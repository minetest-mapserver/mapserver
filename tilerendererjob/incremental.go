package tilerendererjob

import (
	"fmt"
	"mapserver/app"
	"mapserver/settings"
	"mapserver/types"
	"time"

	"github.com/sirupsen/logrus"
)

type IncrementalRenderEvent struct {
	LastMtime int64 `json:"lastmtime"`
}

// returns the position of the last rendered mapblock within the given mtime
// nil if there is none or if it belongs to another mtime (for example a manually changed "last_mtime")
func loadLastMtimePos(s settings.Settings, lastMtime int64) *types.MapBlockCoords {
	var mtime int64
	pos := &types.MapBlockCoords{}

	str := s.GetString(settings.SETTING_LAST_MTIME_POS, "")
	n, err := fmt.Sscanf(str, "%d/%d/%d/%d", &mtime, &pos.X, &pos.Y, &pos.Z)
	if err != nil || n != 4 || mtime != lastMtime {
		return nil
	}

	return pos
}

// saves the mtime and position of the last rendered mapblock
func saveLastMtimePos(s settings.Settings, lastMtime int64, lastPos *types.MapBlockCoords) {
	if lastPos != nil {
		s.SetString(settings.SETTING_LAST_MTIME_POS, fmt.Sprintf("%d/%d/%d/%d", lastMtime, lastPos.X, lastPos.Y, lastPos.Z))
	}
	s.SetInt64(settings.SETTING_LAST_MTIME, lastMtime)
}

func incrementalRender(ctx *app.App) {

	lastMtime := ctx.Settings.GetInt64(settings.SETTING_LAST_MTIME, 0)
	lastPos := loadLastMtimePos(ctx.Settings, lastMtime)

	fields := logrus.Fields{
		"LastMtime": lastMtime,
		"LastPos":   lastPos,
	}
	logrus.WithFields(fields).Info("Starting incremental rendering job")

	for true {
		start := time.Now()

		result, err := ctx.MapBlockAccessor.FindMapBlocksByMtime(lastMtime, lastPos, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			renderingDuration, err := time.ParseDuration(ctx.Config.IncrementalRenderingTimer)
			if err != nil {
				panic(err)
			}
			time.Sleep(renderingDuration)
			continue
		}

		tiles := renderMapblocks(ctx, result.List)

		lastMtime = result.LastMtime
		lastPos = result.LastPos
		saveLastMtimePos(ctx.Settings, lastMtime, lastPos)

		t := time.Now()
		elapsed := t.Sub(start)

		ev := IncrementalRenderEvent{
			LastMtime: result.LastMtime,
		}

		ctx.WebEventbus.Emit("incremental-render-progress", &ev)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
			"elapsed":   elapsed,
			"lastMtime": result.LastMtime,
		}
		logrus.WithFields(fields).Info("incremental rendering")
	}
}
