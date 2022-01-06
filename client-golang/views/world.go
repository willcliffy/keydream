package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/models"
)

type World struct {}

func NewWorld() *World {
	return &World{}
}

func (this *World) Update() (models.State, error) {
	return models.State_WorldConnected, nil
}

func (this *World) Draw(screen *ebiten.Image) {

}

func (this *World) HandleInput() {

}