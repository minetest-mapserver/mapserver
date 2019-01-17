package tilerenderer

import (
	"errors"
	"image"
	"time"
	"bytes"
	"image/png"
	"image/draw"
	"mapserver/coords"
	"mapserver/mapblockrenderer"
	"mapserver/layerconfig"
	"mapserver/tiledb"
	"github.com/sirupsen/logrus"
)

type TileRenderer struct {
	mapblockrenderer *mapblockrenderer.MapBlockRenderer
	layers []layerconfig.Layer
	tdb tiledb.DBAccessor
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer,
	tdb tiledb.DBAccessor,
	layers []layerconfig.Layer) *TileRenderer {

	return &TileRenderer{
		mapblockrenderer: mapblockrenderer,
		layers: layers,
		tdb: tdb,
	}
}

const (
	IMG_SIZE                          = 256
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

	for _, l := range(tr.layers) {
		if l.Id == tc.LayerId {
			layer = &l
		}
	}

	if layer == nil {
		return nil, errors.New("No layer found")
	}

	if tc.Zoom == 13 {
		//max zoomed in on mapblock level
		mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		mbr.Pos1.Y = layer.From
		mbr.Pos2.Y = layer.To

		return tr.mapblockrenderer.Render(mbr.Pos1, mbr.Pos2)
	}

	if tc.Zoom > 13 || tc.Zoom < 1 {
		return nil, errors.New("Invalid zoom")
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
