package game_models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/willcliffy/keydream-server/common"
)

type WorldBroadcast struct {
	ID         common.WorldID `json:"id"`
	IP         string         `json:"ip"`
	NumPlayers int            `json:"num_players"`
}

func BroadcastWorld(wb WorldBroadcast) error {
	payload, err := json.Marshal(wb)
	if err != nil {
		return err
	}

	// TODO: make this configurable. Right now, this is hardcoded to only work with docker compose locally.
	req, err := http.NewRequest("POST", "http://lobby:80/api/v1/worlds", bytes.NewBuffer(payload))
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
