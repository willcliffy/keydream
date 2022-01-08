package objects

import "github.com/willcliffy/keydream/client/common/constants"

type JoinRequest struct {
	WorldID constants.WorldID `json:"id"`
}
