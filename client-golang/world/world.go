package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
)

type World struct {}

func NewWorld(player *common.Player) *World {
	return &World{}
}

func (this *World) Update() (models.State, error) {
	return models.State_WorldConnected, nil
}

func (this *World) Draw(screen *ebiten.Image) {

}

func (this *World) HandleInput() {

}
