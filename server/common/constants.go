package common

import "time"

type PlayerID uint8

type WorldID uint8

const (
	NilPlayerID PlayerID = 0

	// For now, artificially limit the number of players to 20. hypothetical max is 255.
	MaxPlayersPerWorld = 20

	MaxNumberOfWorlds uint8 = 1

	PlayerTimeout = time.Second * 5
)
