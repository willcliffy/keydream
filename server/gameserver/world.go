package gameserver

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/gameserver/models"
)

type World struct {
	ID         common.WorldID
	Players    map[common.PlayerID]*Player
	NumPlayers int
}

func (w World) GetNextPlayerID() common.PlayerID {
	for i := common.PlayerID(1); i <= 255; i++ {
		if w.Players[i] == nil {
			return common.PlayerID(i)
		}
	}

	return common.NilPlayerID
}

func (w *World) ControlLoop(l net.Listener) {
	for {
		if w.NumPlayers == common.MaxPlayersPerWorld {
			time.Sleep(10 * time.Second)
			continue
		}

		c, err := l.Accept()
		if err != nil {
			log.Println("Error connecting:", err.Error())
			return
		}

		newPlayer := NewPlayer(c, w)
		w.Players[newPlayer.ID] = &newPlayer
		go newPlayer.ControlLoop()
		log.Println("Client " + c.RemoteAddr().String() + " connected.")
	}
}

func (w *World) BroadcastLoop() {
	hostname, _ := os.Hostname()

	var failCount uint8
	for {
		// TODO - configure ID, IP. for now, hardcode world 1 and assume no other game servers.
		err := game_models.BroadcastWorld(game_models.WorldBroadcast{
			ID:         1,
			IP:         hostname,
			NumPlayers: w.NumPlayers,
		})

		if err != nil {
			if failCount >= 5 {
				w.Shutdown(true)
				return
			} else {
				log.Printf("Error broadcasting game server: %s\n", err.Error())
				failCount++
			}
		} else {
			failCount = 0
		}

		time.Sleep(10 * time.Second)
	}
}

func (w *World) Initialize() {
	w.Players = make(map[common.PlayerID]*Player)
}

func (w *World) Shutdown(graceful bool) {
	if graceful {
		for {
			if w.NumPlayers == 0 {
				break
			}

			time.Sleep(15 * time.Second)
		}
	} else {
		for _, p := range w.Players {
			if p.Conn != nil {
				p.Conn.Close()
			}
		}
	}
}

func (w *World) OnPlayerJoined(p Player) error {
	log.Printf("Player %d joined the world\n", p.ID)

	if w.NumPlayers >= common.MaxPlayersPerWorld {
		return fmt.Errorf("world is currently full! (capacity %v)", common.MaxPlayersPerWorld)
	}

	if _, ok := w.Players[p.ID]; ok {
		return nil // the player is already in the world
	}

	p.LastUpdated = time.Now()
	w.Players[p.ID] = &p
	w.NumPlayers++

	for _, peer := range w.Players {
		_, err := peer.Conn.Write([]byte(fmt.Sprintf("join %d %d %d\n", p.ID, p.Position.X, p.Position.Y)))
		if err != nil {
			log.Printf("Error sending player join: %s\n", err.Error())
		}
	}

	return nil
}

func (w *World) OnPlayerLeft(p Player) error {
	log.Printf("Player %d left the world\n", p.ID)

	delete(w.Players, p.ID)
	w.NumPlayers--

	for _, p := range w.Players {
		_, err := p.Conn.Write([]byte(fmt.Sprintf("left %d\n", p.ID)))
		if err != nil {
			log.Printf("Error sending player left: %s\n", err.Error())
		}
	}

	return nil
}

func (w *World) OnPlayerMoved(p *Player) error {
	// TODO - validate movement

	if p == nil {
		return fmt.Errorf("player is nil!")
	}

	if _, ok := w.Players[p.ID]; !ok {
		return fmt.Errorf("player not found!")
	}

	p.LastUpdated = time.Now()
	w.Players[p.ID] = p

	for _, p := range w.Players {
		_, err := p.Conn.Write([]byte(fmt.Sprintf("move %d %d %d\n", p.ID, p.Position.X, p.Position.Y)))
		if err != nil {
			log.Printf("Error sending player move: %s\n", err.Error())
		}
	}

	return nil
}
