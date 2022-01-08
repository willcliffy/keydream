package common

import (
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
)


type Player struct {
	ID    constants.PlayerID
	Name  string
	Token string

	World *objects.WorldBroadcast
}

// type Player struct {
// 	Position    game_models.Position
// 	Addr        *net.UDPAddr
// 	LastUpdated time.Time

// 	MessageQueue []string

// 	world *World
// }
