package tilerenderer

import (
	"bytes"
	"errors"
	"image"
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

func resizeImage(src image.Image, tgt image.Image, xoffset int, yoffset int) {
	if src == nil {
		return
	}
	var spix []uint8
	var tpix []uint8

	switch src.(type) {
	case *image.RGBA:
		spix = src.(*image.RGBA).Pix
	case *image.NRGBA:
		spix = src.(*image.NRGBA).Pix
	default:
		log.Print("resizeImage: Got non RGBA non NRGBA source image!")
		return
	}
	switch tgt.(type) {
	case *image.RGBA:
		tpix = tgt.(*image.RGBA).Pix
	case *image.NRGBA:
		tpix = tgt.(*image.NRGBA).Pix
	default:
		log.Print("resizeImage: Got non RGBA non NRGBA target image!")
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
			r := (uint16(spix[six]) + uint16(spix[six+4]) + uint16(spix[six+sinc]) + uint16(spix[six+sinc+4])) >> 2
			g := (uint16(spix[six+1]) + uint16(spix[six+5]) + uint16(spix[six+sinc+1]) + uint16(spix[six+sinc+5])) >> 2
			b := (uint16(spix[six+2]) + uint16(spix[six+6]) + uint16(spix[six+sinc+2]) + uint16(spix[six+sinc+6])) >> 2
			a := (uint16(spix[six+3]) + uint16(spix[six+7]) + uint16(spix[six+sinc+3]) + uint16(spix[six+sinc+7])) >> 2
			tpix[tix] = uint8(r)
			tpix[tix+1] = uint8(g)
			tpix[tix+2] = uint8(b)
			tpix[tix+3] = uint8(a)
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
	err := tr.renderImage(tc)

	if err != nil {
		return err
	}

	return nil
}

func (tr *TileRenderer) getCachedImage(tc *coords.TileCoords) (image.Image, error) {

	cachedtile, err := tr.tdb.GetTile(tc)
	if err != nil {
		return nil, err
	}
	if cachedtile == nil {
		return nil, nil
	}

	reader := bytes.NewReader(cachedtile)
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (tr *TileRenderer) renderImage(tc *coords.TileCoords) (error) {
	if tc.Zoom < coords.MIN_ZOOM || tc.Zoom > coords.MAX_ZOOM {
		return errors.New("Invalid zoom")
	}

	// Image existing and not outdated --> Done
	if ! tr.tdb.IsOutdated(tc) && tr.tdb.TileExists(tc) {
		return nil
	}

	// No cached image (or outdated)
	currentLayer := layer.FindLayerById(tr.layers, tc.LayerId)

	if currentLayer == nil {
		return errors.New("No layer found")
	}

	renderedTiles.With(prometheus.Labels{"zoom": strconv.Itoa(tc.Zoom)}).Inc()
	timer := prometheus.NewTimer(renderDuration)
	defer timer.ObserveDuration()

	// Origin zoom
	if tc.Zoom == coords.MAX_ZOOM {
		log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("RenderMapblock")

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

			return err
		}

		if img == nil {
			return nil
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		tile := buf.Bytes()
		tr.tdb.SetTile(tc, tile)

		tr.Eventbus.Emit(eventbus.TILE_RENDERED, tc)

		return nil

	} else {
		// Unzoomed

		log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("RenderImage")

		quads := tc.GetZoomedQuadrantsFromTile()

		fields := logrus.Fields{
			"UpperLeft":  quads.UpperLeft,
			"UpperRight": quads.UpperRight,
			"LowerLeft":  quads.LowerLeft,
			"LowerRight": quads.LowerRight,
		}
		log.WithFields(fields).Debug("Quad image stats")

		start := time.Now()

		upperLeft, err := tr.getCachedImage(quads.UpperLeft)
		if err != nil {
			return err
		}

		upperRight, err := tr.getCachedImage(quads.UpperRight)
		if err != nil {
			return err
		}

		lowerLeft, err := tr.getCachedImage(quads.LowerLeft)
		if err != nil {
			return err
		}

		lowerRight, err := tr.getCachedImage(quads.LowerRight)
		if err != nil {
			return err
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

		return nil
	}
}
