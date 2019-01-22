package tilerenderer

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/png"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layer"
	"mapserver/mapblockrenderer"
	"mapserver/mapobjectdb"
	"time"

	"github.com/disintegration/imaging"
	"github.com/sirupsen/logrus"
)

type TileRenderer struct {
	mapblockrenderer *mapblockrenderer.MapBlockRenderer
	layers           []layer.Layer
	tdb              mapobjectdb.DBAccessor
	dba              db.DBAccessor
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer,
	tdb mapobjectdb.DBAccessor,
	dba db.DBAccessor,
	layers []layer.Layer) *TileRenderer {

	return &TileRenderer{
		mapblockrenderer: mapblockrenderer,
		layers:           layers,
		tdb:              tdb,
		dba:              dba,
	}
}

const (
	IMG_SIZE = 256
)

func (tr *TileRenderer) Render(tc *coords.TileCoords, recursionDepth int) ([]byte, error) {

	//Check cache
	tile, err := tr.tdb.GetTile(tc)
	if err != nil {
		return nil, err
	}

	if tile == nil {

		if recursionDepth == 0 {
			log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("Skip image")
			return nil, nil
		}

		//No tile in db
		img, data, err := tr.RenderImage(tc, recursionDepth)

		if err != nil {
			return nil, err
		}

		if img == nil {
			//empty tile
			return nil, nil
		}

		return data, nil
	}

	return tile.Data, nil
}

func (tr *TileRenderer) RenderImage(tc *coords.TileCoords, recursionDepth int) (*image.NRGBA, []byte, error) {

	cachedtile, err := tr.tdb.GetTile(tc)
	if err != nil {
		return nil, nil, err
	}

	if cachedtile != nil {
		reader := bytes.NewReader(cachedtile.Data)
		cachedimg, err := png.Decode(reader)
		if err != nil {
			return nil, nil, err
		}

		rect := image.Rectangle{
			image.Point{0, 0},
			image.Point{IMG_SIZE, IMG_SIZE},
		}

		img := image.NewNRGBA(rect)
		draw.Draw(img, rect, cachedimg, image.ZP, draw.Src)

		log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("Cached image")
		return img, cachedtile.Data, nil
	}

	if recursionDepth == 0 {
		log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("Skip image")
		return nil, nil, nil
	}

	log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("RenderImage")

	var layer *layer.Layer

	for _, l := range tr.layers {
		if l.Id == tc.LayerId {
			layer = &l
		}
	}

	if layer == nil {
		return nil, nil, errors.New("No layer found")
	}

	if tc.Zoom > 13 || tc.Zoom < 1 {
		return nil, nil, errors.New("Invalid zoom")
	}

	if tc.Zoom == 13 {
		//max zoomed in on mapblock level
		mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		mbr.Pos1.Y = layer.From
		mbr.Pos2.Y = layer.To

		img, err := tr.mapblockrenderer.Render(mbr.Pos1, mbr.Pos2)

		if err != nil {
			return nil, nil, err
		}

		if img == nil {
			return nil, nil, nil
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		return img, buf.Bytes(), nil
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

	upperLeft, _, err := tr.RenderImage(quads.UpperLeft, recursionDepth-1)
	if err != nil {
		return nil, nil, err
	}

	upperRight, _, err := tr.RenderImage(quads.UpperRight, recursionDepth-1)
	if err != nil {
		return nil, nil, err
	}

	lowerLeft, _, err := tr.RenderImage(quads.LowerLeft, recursionDepth-1)
	if err != nil {
		return nil, nil, err
	}

	lowerRight, _, err := tr.RenderImage(quads.LowerRight, recursionDepth-1)
	if err != nil {
		return nil, nil, err
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

	rect := image.Rect(0, 0, 128, 128)
	if upperLeft != nil {
		resizedImg := imaging.Resize(upperLeft, 128, 128, imaging.Lanczos)
		draw.Draw(img, rect, resizedImg, image.ZP, draw.Src)
	}

	rect = image.Rect(128, 0, 256, 128)
	if upperRight != nil {
		resizedImg := imaging.Resize(upperRight, 128, 128, imaging.Lanczos)
		draw.Draw(img, rect, resizedImg, image.ZP, draw.Src)
	}

	rect = image.Rect(0, 128, 128, 256)
	if lowerLeft != nil {
		resizedImg := imaging.Resize(lowerLeft, 128, 128, imaging.Lanczos)
		draw.Draw(img, rect, resizedImg, image.ZP, draw.Src)
	}

	rect = image.Rect(128, 128, 256, 256)
	if lowerRight != nil {
		resizedImg := imaging.Resize(lowerRight, 128, 128, imaging.Lanczos)
		draw.Draw(img, rect, resizedImg, image.ZP, draw.Src)
	}

	t = time.Now()
	quadresize := t.Sub(start)
	start = t

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	t = time.Now()
	encode := t.Sub(start)
	start = t

	tile := mapobjectdb.Tile{Pos: tc, Data: buf.Bytes(), Mtime: time.Now().Unix()}
	tr.tdb.SetTile(&tile)

	t = time.Now()
	cache := t.Sub(start)

	fields = logrus.Fields{
		"X":          tc.X,
		"Y":          tc.Y,
		"Zoom":       tc.Zoom,
		"LayerId":    tc.LayerId,
		"size": len(tile.Data),
		"quadrender": quadrender,
		"quadresize": quadresize,
		"encode":     encode,
		"cache":      cache,
	}
	log.WithFields(fields).Debug("Cross stitch")

	return img, buf.Bytes(), nil
}
