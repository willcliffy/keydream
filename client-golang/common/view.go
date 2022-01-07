package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/models"
)

type View interface {
	//Layout(outsideWidth, outsideHeight int) (int, int)
	Update() (models.State, error)
	Draw(screen *ebiten.Image)
}
