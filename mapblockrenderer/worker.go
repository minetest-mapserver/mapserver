package mapblockrenderer

import (
  "time"
  "bytes"
  "mapserver/coords"
  "image/png"

)

type JobData struct {
	Pos1, Pos2 coords.MapBlockCoords
	X, Z       int
}

type JobResult struct {
  Data *bytes.Buffer
  Duration time.Duration
  Job JobData
}

func worker(r *MapBlockRenderer, jobs <-chan JobData, results chan<- JobResult) {
	for d := range jobs {
		img, _ := r.Render(d.Pos1, d.Pos2)

    w := new(bytes.Buffer)
    start := time.Now()

		if img != nil {
			png.Encode(w, img)
		}

    t := time.Now()
    elapsed := t.Sub(start)

    res := JobResult{Data: w, Duration: elapsed, Job: d}
    results <- res

	}
}
