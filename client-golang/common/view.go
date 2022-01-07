package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/models"
)

type KeydreamView interface {
	Update() (models.State, error)
	Draw(screen *ebiten.Image)
}
