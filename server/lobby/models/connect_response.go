package lobby_models

import (
	game_models "github.com/willcliffy/keydream/server/world/models"
)

type ConnectResponse struct {
	Worlds []game_models.WorldBroadcast
}
