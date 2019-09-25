package tilerenderer

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/png"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/eventbus"
	"mapserver/layer"
	"mapserver/mapblockrenderer"
	"mapserver/tiledb"
	"strconv"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

type TileRenderer struct {
	mapblockrenderer *mapblockrenderer.MapBlockRenderer
	layers           []*layer.Layer
	tdb              *tiledb.TileDB
	dba              db.DBAccessor
	Eventbus         *eventbus.Eventbus
}

func resizeImage(src *image.NRGBA, tgt *image.NRGBA, xoffset int, yoffset int) {
	if src == nil {
		return
	}
	w := src.Bounds().Dy() >> 1
	h := src.Bounds().Dx() >> 1
	sinc := src.Bounds().Dy() * 4
	tinc := tgt.Bounds().Dx() * 4

	for y := 0; y < h; y++ {
		six := y * sinc * 2
		tix := 4 * xoffset + (yoffset + y) * tinc
		for x := 0; x < w; x++ {
			r := (uint16(src.Pix[six]) + uint16(src.Pix[six+4]) + uint16(src.Pix[six+sinc]) + uint16(src.Pix[six+sinc+4])) >> 2
			g := (uint16(src.Pix[six+1]) + uint16(src.Pix[six+5]) + uint16(src.Pix[six+sinc+1]) + uint16(src.Pix[six+sinc+5])) >> 2
			b := (uint16(src.Pix[six+2]) + uint16(src.Pix[six+6]) + uint16(src.Pix[six+sinc+2]) + uint16(src.Pix[six+sinc+6])) >> 2
			a := (uint16(src.Pix[six+3]) + uint16(src.Pix[six+7]) + uint16(src.Pix[six+sinc+3]) + uint16(src.Pix[six+sinc+7])) >> 2
			tgt.Pix[tix] = uint8(r)
			tgt.Pix[tix+1] = uint8(g)
			tgt.Pix[tix+2] = uint8(b)
			tgt.Pix[tix+3] = uint8(a)
			tix+=4
			six+=8
		}
	}
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer,
	tdb *tiledb.TileDB,
	dba db.DBAccessor,
	layers []*layer.Layer) *TileRenderer {

	return &TileRenderer{
		mapblockrenderer: mapblockrenderer,
		layers:           layers,
		tdb:              tdb,
		dba:              dba,
		Eventbus:         eventbus.New(),
	}
}

const (
	IMG_SIZE = 256
	SUB_IMG_SIZE = IMG_SIZE >> 1
)

func (tr *TileRenderer) Render(tc *coords.TileCoords) (error) {
	//No tile in db
	_, err := tr.renderImage(tc, 2)

	if err != nil {
		return err
	}

	return nil
}

func (tr *TileRenderer) renderImage(tc *coords.TileCoords, recursionDepth int) (*image.NRGBA, error) {

	if recursionDepth < 2 {
		cachedtile, err := tr.tdb.GetTile(tc)
		if err != nil {
			return nil, err
		}

		if cachedtile != nil {
			reader := bytes.NewReader(cachedtile)
			cachedimg, err := png.Decode(reader)
			if err != nil {
				return nil, err
			}

			rect := image.Rectangle{
				image.Point{0, 0},
				image.Point{IMG_SIZE, IMG_SIZE},
			}

			img := image.NewNRGBA(rect)
			draw.Draw(img, rect, cachedimg, image.ZP, draw.Src)

			log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("Cached image")
			return img, nil
		}
	}

	if recursionDepth <= 1 && tc.Zoom < 13 {
		//non-cached layer and not in "origin" zoom, skip tile
		log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("Skip image")
		return nil, nil
	}

	log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("RenderImage")

	renderedTiles.With(prometheus.Labels{"zoom": strconv.Itoa(tc.Zoom)}).Inc()
	timer := prometheus.NewTimer(renderDuration)
	defer timer.ObserveDuration()

	currentLayer := layer.FindLayerById(tr.layers, tc.LayerId)

	if currentLayer == nil {
		return nil, errors.New("No layer found")
	}

	if tc.Zoom > 13 || tc.Zoom < 1 {
		return nil, errors.New("Invalid zoom")
	}

	if tc.Zoom == 13 {
		//max zoomed in on mapblock level
		mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		mbr.Pos1.Y = currentLayer.From
		mbr.Pos2.Y = currentLayer.To

		img, err := tr.mapblockrenderer.Render(mbr.Pos1, mbr.Pos2)

		if err != nil {
			fields := logrus.Fields{
				"pos1": mbr.Pos1,
				"pos2": mbr.Pos2,
				"err":  err,
			}
			log.WithFields(fields).Debug("mapblock render from tilerender")

			return nil, err
		}

		if img == nil {
			return nil, nil
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		tile := buf.Bytes()
		tr.tdb.SetTile(tc, tile)

		return img, nil
	}

	//zoom 1-12
	quads := tc.GetZoomedQuadrantsFromTile()

	fields := logrus.Fields{
		"UpperLeft":  quads.UpperLeft,
		"UpperRight": quads.UpperRight,
		"LowerLeft":  quads.LowerLeft,
		"LowerRight": quads.LowerRight,
	}
	log.WithFields(fields).Debug("Quad image stats")

	start := time.Now()

	upperLeft, err := tr.renderImage(quads.UpperLeft, recursionDepth-1)
	if err != nil {
		return nil, err
	}

	upperRight, err := tr.renderImage(quads.UpperRight, recursionDepth-1)
	if err != nil {
		return nil, err
	}

	lowerLeft, err := tr.renderImage(quads.LowerLeft, recursionDepth-1)
	if err != nil {
		return nil, err
	}

	lowerRight, err := tr.renderImage(quads.LowerRight, recursionDepth-1)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	quadrender := t.Sub(start)
	start = t

	img := image.NewNRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{IMG_SIZE, IMG_SIZE},
		},
	)

	resizeImage(upperLeft, img, 0, 0)
	resizeImage(upperRight, img, SUB_IMG_SIZE, 0)
	resizeImage(lowerLeft, img, 0, SUB_IMG_SIZE)
	resizeImage(lowerRight, img, SUB_IMG_SIZE, SUB_IMG_SIZE)

	t = time.Now()
	quadresize := t.Sub(start)
	start = t

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	t = time.Now()
	encode := t.Sub(start)
	start = t

	tile := buf.Bytes()
	tr.tdb.SetTile(tc, tile)

	t = time.Now()
	cache := t.Sub(start)

	fields = logrus.Fields{
		"X":          tc.X,
		"Y":          tc.Y,
		"Zoom":       tc.Zoom,
		"LayerId":    tc.LayerId,
		"size":       len(tile),
		"quadrender": quadrender,
		"quadresize": quadresize,
		"encode":     encode,
		"cache":      cache,
	}
	log.WithFields(fields).Debug("Cross stitch")

	tr.Eventbus.Emit(eventbus.TILE_RENDERED, tc)

	return img, nil
}
