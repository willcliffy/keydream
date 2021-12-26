package lobby_models

import (
	game_models "github.com/willcliffy/keydream-server/gameserver/models"
)

type ConnectResponse struct {
	Worlds []game_models.WorldBroadcast
}
