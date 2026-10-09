package mapblockrenderer

import (
	"bytes"
	"image/png"
	"mapserver/types"
	"time"

	"github.com/sirupsen/logrus"
)

type JobData struct {
	Pos1, Pos2 *types.MapBlockCoords
}

type JobResult struct {
	Data     *bytes.Buffer
	Duration time.Duration
	Job      JobData
}

func Worker(r *MapBlockRenderer, jobs <-chan JobData, results chan<- JobResult) {
	for d := range jobs {
		img, err := r.Render(d.Pos1, d.Pos2)
		if err != nil {
			log.WithFields(logrus.Fields{"err": err}).Error("render")
		}

		w := new(bytes.Buffer)
		start := time.Now()

		if img != nil {
			if err := png.Encode(w, img); err != nil {
				log.WithFields(logrus.Fields{"err": err}).Error("png-encode")
				w.Reset()
			}
		}

		t := time.Now()
		elapsed := t.Sub(start)

		res := JobResult{Data: w, Duration: elapsed, Job: d}
		results <- res

	}
}
