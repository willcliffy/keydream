package world

import (
	"net"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/world/models"
)

type Player struct {
	ID          common.PlayerID
	Name        string
	Position    game_models.Position
	Conn        net.PacketConn
	LastUpdated time.Time

	world *World
}

func NewPlayer(world *World, name string) Player {
	return Player{
		ID:          world.GetNextPlayerID(),
		Name:        name,
		Position:    game_models.NewPosition(),
		LastUpdated: time.Now(),

		world: world,
	}
}
