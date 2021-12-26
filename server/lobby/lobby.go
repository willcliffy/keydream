package lobby

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/world/models"
)

type LobbyHandler struct {
	shutdown chan bool
	Worlds   map[common.WorldID]game_models.WorldBroadcast
}

func (l *LobbyHandler) ControlLoop() {
	for {
		select {
		case <-l.shutdown:
			return
		case <-time.After(time.Second * 5):
			var s string
			for _, world := range l.Worlds {
				s += fmt.Sprintf("\n\t%d: %s (%v)", world.ID, world.IP, world.NumPlayers)
			}

			log.Printf("Game server list: %+v\n", s)
			time.Sleep(10 * time.Second)
		}
	}
}

func (l *LobbyHandler) UpdateWorldHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes := new(bytes.Buffer)
	_, err := bodyBytes.ReadFrom(r.Body)

	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	var world game_models.WorldBroadcast

	if err = json.Unmarshal(bodyBytes.Bytes(), &world); err != nil {
		http.Error(w, "unable to unmarshal request body", http.StatusBadRequest)
		return
	}

	l.Worlds[world.ID] = world

	common.SendHTTPResponse(w, nil, http.StatusNoContent)
}
