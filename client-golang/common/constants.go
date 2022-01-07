package common

const (
	TileSize = 16
	TileScale = 2

	TileSizeScaled = TileSize * TileScale

	ScreenWidthInTiles = 25
	ScreenHeightInTiles = 15

	ScreenWidth = TileSize * TileScale * ScreenWidthInTiles
	ScreenHeight = TileSize * TileScale * ScreenHeightInTiles

	DefaultButtonWidth = TileSizeScaled * 5
	DefaultButtonHeight = TileSizeScaled * 1.5
)
