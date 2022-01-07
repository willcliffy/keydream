package common

import "time"

// Client-specific constants

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

// Constants from server

type PlayerID uint8

type WorldID uint8

const (
	NilPlayerID PlayerID = 0

	// For now, artificially limit the number of players to 20. hypothetical max is 255.
	MaxPlayersPerWorld = 20

	MaxNumberOfWorlds uint8 = 1

	PlayerTimeout = time.Second * 5
)
