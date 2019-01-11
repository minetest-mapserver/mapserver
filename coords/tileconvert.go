package coords

const (
	MAX_ZOOM = 13
)


func GetTileCoordsFromMapBlock(mbc MapBlockCoords) TileCoords {
	return TileCoords{X:mbc.X, Y:(mbc.Z + 1) * -1, Zoom:MAX_ZOOM};
}
