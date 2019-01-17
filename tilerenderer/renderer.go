package tilerenderer

import (
	"image"
	"mapserver/coords"
	"mapserver/mapblockrenderer"
)

type TileRenderer struct {
	mapblockrenderer *mapblockrenderer.MapBlockRenderer
}

func NewTileRenderer(mapblockrenderer *mapblockrenderer.MapBlockRenderer) *TileRenderer {
	return &TileRenderer{
		mapblockrenderer: mapblockrenderer,
	}
}

//TODO layerConfig
func (tr *TileRenderer) Render(tc coords.TileCoords) (*image.NRGBA, error) {
	if tc.Zoom == 13 {
		//max zoomed in on mapblock level
		//mbr := coords.GetMapBlockRangeFromTile(tc, 0)
		//return tr.mapblockrenderer.Render()

	}

	return nil, nil
}
