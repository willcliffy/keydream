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
	Conn       *net.UDPConn
}

func NewWorld(id common.WorldID, udpConn *net.UDPConn) *World {
	return &World{
		ID:         id,
		Players:    make(map[common.PlayerID]*Player),
		NumPlayers: 0,
		Conn:       udpConn,
	}
}

func (w World) GetNextPlayerID() common.PlayerID {
	for i := common.PlayerID(1); i <= 255; i++ {
		if w.Players[i] == nil {
			return common.PlayerID(i)
		} else if w.Players[i].LastUpdated.Before(time.Now().Add(-common.PlayerTimeout)) {
			w.OnPlayerLeft(*w.Players[i])
			return common.PlayerID(i)
		}
	}

	return common.NilPlayerID
}

func (w *World) ControlLoop() {
	for {
		buf := make([]byte, 128)
		_, addr, err := w.Conn.ReadFromUDP(buf)
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
			playerID, err := strconv.Atoi(msg[1])
			if err != nil {
				log.Printf("Error parsing player ID: %s\n", err.Error())
				continue
			}

			player, ok := w.Players[common.PlayerID(playerID)]
			if !ok {
				log.Printf("Player %d not found\n", playerID)
				continue
			}

			player.OnTick()
		case "join":
			if len(msg) != 3 {
				log.Printf("Invalid join message: %s", string(buf))
				continue
			}

			if w.NumPlayers >= common.MaxPlayersPerWorld {
				log.Printf("World is full!\n")
				time.Sleep(10 * time.Second)
				continue
			}

			// todo - validate player name
			playerName := msg[1]

			w.OnPlayerJoined(playerName, addr)
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

			if err = w.OnPlayerMoved(player); err != nil {
				log.Printf("Error moving player: %s\n", err.Error())
				continue
			}

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
		for _, peer := range w.Players {
			_, _ = w.Conn.WriteToUDP([]byte("disconnect \n"), peer.Addr)
		}
	}
}

func (w *World) OnPlayerJoined(name string, addr *net.UDPAddr) {
	p := Player{
		ID:          w.GetNextPlayerID(),
		Name:        name,
		Addr:        addr,
		Position:    game_models.NewPosition(),
		LastUpdated: time.Now(),

		world: w,
	}

	w.Players[p.ID] = &p

	for _, peer := range w.Players {
		peer.OnPlayerJoined(&p)
	}

	_, err := w.Conn.WriteToUDP([]byte(fmt.Sprintf("%d \n", p.ID)), addr)
	if err != nil {
		log.Printf("Error sending player id: %s\n", err.Error())
	}

	log.Printf("Player %d joined the world\n", p.ID)

	p.OnTick()
}

func (w *World) OnPlayerLeft(p Player) {
	if w.Players[p.ID] != nil {
		w.NumPlayers--
	}

	delete(w.Players, p.ID)

	for _, peer := range w.Players {
		peer.OnPlayerLeft(&p)
	}

	log.Printf("Player %d left the world\n", p.ID)
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

	for _, peer := range w.Players {
		peer.OnPlayerMoved(p)
	}

	p.OnTick()

	return nil
}
