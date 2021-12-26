package gameserver

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/gameserver/models"
)

type Player struct {
	ID          common.PlayerID
	Name        string
	Position    game_models.Position
	Conn        net.Conn
	LastUpdated time.Time

	world  *World
}

func NewPlayer(conn net.Conn, world *World) Player {
	return Player{
		ID:          world.GetNextPlayerID(),
		Name:        "",
		Position:    game_models.NewPosition(),
		Conn:        conn,
		LastUpdated: time.Now(),

		world:       world,
	}
}

func (p Player) ControlLoop() {
	if err := p.world.OnPlayerJoined(p); err != nil {
		log.Printf("Error handling player join: %s\n", err.Error())
		_, _ = p.Conn.Write([]byte("error\n"))
	}

	for {
		if p.Conn == nil {
			log.Printf("Player %s disconnected.", p.Name)
			_ = p.world.OnPlayerLeft(p)
			return
		}

		buffer, err := bufio.NewReader(p.Conn).ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Player %s disconnected.", p.Name)
			} else {
				log.Println("Error reading player message:", err.Error())
			}
			break
		}

		// TODO - handle player commands. for now just echo back to client
		log.Printf("Echoing player message: %s\n", string(buffer[:len(buffer)-1]))
		if _, err = p.Conn.Write(buffer); err != nil {
			log.Println("Error writing player message:", err.Error())
			break
		}
	}

	p.Conn.Close()
}
