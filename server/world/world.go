package world

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/world/models"
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

func (w *World) ControlLoop(l *net.UDPConn) {
	for {
		buf := make([]byte, 128)
		_, addr, err := l.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading from packet conn: %s\n", err.Error())
			continue
		}

		if len(buf) == 0 {
			continue
		}

		msg := strings.Split(string(buf), " ")

		switch msg[0] {
		case "tick":
			_, _ = l.WriteToUDP([]byte("tock\n"), addr)
		case "join":
			if len(msg) != 2 {
				log.Printf("Invalid join message: %s", string(buf))
				continue
			}

			if w.NumPlayers == common.MaxPlayersPerWorld {
				log.Printf("World is full!\n")
				time.Sleep(10 * time.Second)
				continue
			}

			playerName := msg[1]

			log.Printf("Client %v connected: %v", addr, string(buf))
			newPlayer := NewPlayer(w, playerName)
			w.Players[newPlayer.ID] = &newPlayer
			_, err = l.WriteToUDP([]byte(fmt.Sprintf("%d\n", newPlayer.ID)), addr)
			if err != nil {
				log.Printf("Error sending player id: %s\n", err.Error())
			}
		case "move":
			if len(msg) != 5 {
				log.Printf("Invalid move message: '%s' of len %v\n", string(buf), len(msg))
				continue
			}

			playerID, err := strconv.Atoi(msg[1])
			if err != nil {
				log.Printf("Error parsing player id: %s\n", err.Error())
				continue
			}

			player, ok := w.Players[common.PlayerID(playerID)]

			if !ok {
				log.Printf("Player %d not found\n", playerID)
				continue
			}

			x, err := strconv.ParseInt(msg[2], 10, 64)
			if err != nil {
				log.Printf("Error parsing x: %s\n", err.Error())
				continue
			}

			y, err := strconv.ParseInt(msg[3], 10, 64)
			if err != nil {
				log.Printf("Error parsing y: %s\n", err.Error())
				continue
			}

			player.Position.X = x
			player.Position.Y = y

			_ = w.OnPlayerMoved(player)

			log.Printf("Player %d moved to %d, %d\n", playerID, x, y)
		default:
			log.Printf("Unknown message: '%v'", msg[0])
			continue
		}
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

	// for _, peer := range w.Players {
	// 	_, err := peer.Conn.Write([]byte(fmt.Sprintf("join %d %d %d\n", p.ID, p.Position.X, p.Position.Y)))
	// 	if err != nil {
	// 		log.Printf("Error sending player join: %s\n", err.Error())
	// 	}
	// }

	return nil
}

func (w *World) OnPlayerLeft(p Player) error {
	log.Printf("Player %d left the world\n", p.ID)

	if w.Players[p.ID] != nil {
		w.NumPlayers--
	}

	delete(w.Players, p.ID)

	// for _, p := range w.Players {
	// 	_, err := p.Conn.Write([]byte(fmt.Sprintf("left %d\n", p.ID)))
	// 	if err != nil {
	// 		log.Printf("Error sending player left: %s\n", err.Error())
	// 	}
	// }

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

	// for _, p := range w.Players {
	// 	_, err := p.Conn.Write([]byte(fmt.Sprintf("move %d %d %d\n", p.ID, p.Position.X, p.Position.Y)))
	// 	if err != nil {
	// 		log.Printf("Error sending player move: %s\n", err.Error())
	// 	}
	// }

	return nil
}
