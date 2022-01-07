package objects

import "github.com/willcliffy/keydream/client/common"

type JoinRequest struct {
	WorldID common.WorldID `json:"id"`
}
