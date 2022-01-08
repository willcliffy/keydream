package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/willcliffy/keydream/client/common/constants"
)

type WorldBroadcast struct {
	ID          constants.WorldID `json:"id"`
	IP          string         `json:"ip"`
	NumPlayers  int            `json:"num_players"`
	LastUpdated time.Time      `json:"-"`
}

func BroadcastWorld(wb WorldBroadcast) error {
	payload, err := json.Marshal(wb)
	if err != nil {
		return err
	}

	// TODO: make this configurable. Right now, this is hardcoded to only work with docker compose locally.
	req, err := http.NewRequest("POST", "http://lobby.keydream.tk/api/v1/worlds", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response status from lobby: %d", res.StatusCode)
	}

	return nil
}
