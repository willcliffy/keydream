package world

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/world/models"
)

type Player struct {
	ID          common.PlayerID
	Name        string
	Position    game_models.Position
	Addr        *net.UDPAddr
	LastUpdated time.Time

	MessageQueue []string

	world *World
}

func NewPlayer(world *World, name string, addr *net.UDPAddr) Player {
	return Player{
		ID:          world.GetNextPlayerID(),
		Name:        name,
		Position:    game_models.NewPosition(),
		Addr:        addr,
		LastUpdated: time.Now(),

		world: world,
	}
}

func (p *Player) OnTick() {
	if len(p.MessageQueue) == 0 {
		_, err := p.world.Conn.WriteToUDP([]byte("tock \n"), p.Addr)
		if err != nil {
			log.Printf("Error sending tock: %s\n", err.Error())
		}
	} else {
		for _, msg := range p.MessageQueue {
			_, err := p.world.Conn.WriteToUDP([]byte(msg+"\n"), p.Addr)
			if err != nil {
				log.Printf("Error sending message to client: %s\n", err.Error())
			}
		}

		p.MessageQueue = []string{}
	}

	p.LastUpdated = time.Now()
}

func (p *Player) OnPlayerJoined(oth *Player) {
	p.MessageQueue = append(p.MessageQueue, fmt.Sprintf("join %d %s %d %d", oth.ID, oth.Name, oth.Position.X, oth.Position.Y))
}

func (p *Player) OnPlayerLeft(oth *Player) {
	p.MessageQueue = append(p.MessageQueue, fmt.Sprintf("left %d", oth.ID))
}

func (p *Player) OnPlayerMoved(oth *Player) {
	p.MessageQueue = append(p.MessageQueue, fmt.Sprintf("move %d %d %d", oth.ID, oth.Position.X, oth.Position.Y))
}
