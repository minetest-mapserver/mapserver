package tilerenderer

import (
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"image"
	"image/draw"
	"image/png"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layerconfig"
	"mapserver/mapblockrenderer"
	"mapserver/tiledb"
	"time"
)

type TileRenderer struct {
	mapblockrenderer *mapblockrenderer.MapBlockRenderer
	layers           []layerconfig.Layer
	tdb              tiledb.DBAccessor
	dba              db.DBAccessor
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer,
	tdb tiledb.DBAccessor,
	dba db.DBAccessor,
	layers []layerconfig.Layer) *TileRenderer {

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

func (tr *TileRenderer) Render(tc coords.TileCoords) ([]byte, error) {

	//Check cache
	tile, err := tr.tdb.GetTile(tc)
	if err != nil {
		return nil, err
	}

	if tile == nil {
		//No tile in db
		img, err := tr.RenderImage(tc)

		if err != nil {
			return nil, err
		}

		if img == nil {
			//empty tile
			return nil, nil
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		return buf.Bytes(), nil
	}

	return tile.Data, nil
}

func (tr *TileRenderer) RenderImage(tc coords.TileCoords) (*image.NRGBA, error) {

	cachedtile, err := tr.tdb.GetTile(tc)
	if err != nil {
		return nil, err
	}

	if cachedtile != nil {
		reader := bytes.NewReader(cachedtile.Data)
		cachedimg, err := png.Decode(reader)
		if err != nil {
			return nil, err
		}

		return cachedimg.(*image.NRGBA), nil
	}

	log.WithFields(logrus.Fields{"x": tc.X, "y": tc.Y, "zoom": tc.Zoom}).Debug("RenderImage")

	var layer *layerconfig.Layer

	for _, l := range tr.layers {
		if l.Id == tc.LayerId {
			layer = &l
		}
	}

	if layer == nil {
		return nil, errors.New("No layer found")
	}

	if tc.Zoom > 13 || tc.Zoom < 1 {
		return nil, errors.New("Invalid zoom")
	}

	if tc.Zoom == 13 {
		//max zoomed in on mapblock level
		mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		mbr.Pos1.Y = layer.From
		mbr.Pos2.Y = layer.To

		//count blocks
		count, err := tr.dba.CountBlocks(mbr.Pos1, mbr.Pos2)

		if err != nil {
			return nil, err
		}

		if count == 0 {
			return nil, nil
		}

		return tr.mapblockrenderer.Render(mbr.Pos1, mbr.Pos2)
	}

	if tc.Zoom == 12 {
		//count blocks and ignore empty tiles
		mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		mbr.Pos1.Y = layer.From
		mbr.Pos2.Y = layer.To

		count, err := tr.dba.CountBlocks(mbr.Pos1, mbr.Pos2)

		if err != nil {
			return nil, err
		}

		if count == 0 {
			return nil, nil
		}
	}

	//zoom 1-12
	quads := tc.GetZoomedQuadrantsFromTile()

	upperLeft, err := tr.RenderImage(quads.UpperLeft)
	if err != nil {
		return nil, err
	}

	upperRight, err := tr.RenderImage(quads.UpperRight)
	if err != nil {
		return nil, err
	}

	lowerLeft, err := tr.RenderImage(quads.LowerLeft)
	if err != nil {
		return nil, err
	}

	lowerRight, err := tr.RenderImage(quads.LowerRight)
	if err != nil {
		return nil, err
	}

	isEmpty := upperLeft == nil && upperRight == nil && lowerLeft == nil && lowerRight == nil

	if isEmpty && (tc.Zoom == 11 || tc.Zoom == 10) {
		//don't cache empty zoomed tiles
		return nil, nil
	}

	img := image.NewNRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{IMG_SIZE, IMG_SIZE},
		},
	)

	rect := image.Rect(0, 0, 128, 128)
	if upperLeft != nil {
		draw.Draw(img, rect, upperLeft, image.ZP, draw.Src)
	}

	rect = image.Rect(128, 0, 256, 128)
	if upperRight != nil {
		draw.Draw(img, rect, upperRight, image.ZP, draw.Src)
	}

	rect = image.Rect(0, 128, 128, 256)
	if lowerLeft != nil {
		draw.Draw(img, rect, lowerLeft, image.ZP, draw.Src)
	}

	rect = image.Rect(128, 128, 256, 256)
	if lowerRight != nil {
		draw.Draw(img, rect, lowerRight, image.ZP, draw.Src)
	}

	buf := new(bytes.Buffer)
	if img != nil {
		png.Encode(buf, img)
	}

	tile := tiledb.Tile{Pos: tc, Data: buf.Bytes(), Mtime: time.Now().Unix()}
	tr.tdb.SetTile(&tile)

	return img, nil
}
