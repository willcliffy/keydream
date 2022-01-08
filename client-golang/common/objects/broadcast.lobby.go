package objects

import (
	"time"

	"github.com/willcliffy/keydream/client/common/constants"
)

type WorldBroadcast struct {
	ID          constants.WorldID `json:"id"`
	IP          string         `json:"ip"`
	NumPlayers  int            `json:"num_players"`
	LastUpdated time.Time      `json:"-"`
}
