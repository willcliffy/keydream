package lobby_models

import "github.com/willcliffy/keydream/server/common"

type JoinRequest struct {
	WorldID common.WorldID `json:"id"`
}
