package constants

import "time"

// Client-specific constants

const (
	TileSize = 16
	TileScale = 2
	CharacterScale = 2

	TileSizeScaled = TileSize * TileScale

	ScreenWidthInTiles = 25
	ScreenHeightInTiles = 15

	ScreenWidth = TileSize * TileScale * ScreenWidthInTiles
	ScreenHeight = TileSize * TileScale * ScreenHeightInTiles

	DefaultButtonWidth = TileSizeScaled * 5
	DefaultButtonHeight = TileSizeScaled * 1.5

	// for now, quarter of a second
	CharacterAnimationSpeed = 250 * time.Millisecond
	CharacterWalkSpeed float64 = 2
	RemoteCharacterMinWalkSpeed float64 = 0.5
	RemoteCharacterMaxWalkSpeed float64 = 0.99 * CharacterWalkSpeed
	RemoteCharacterWalkAcceleration = 0.001

	RemoteCharacterAlpha = 1.025 * RemoteCharacterMaxWalkSpeed
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

	WorldTickRate = 250 * time.Millisecond
)
