package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/models"
)

type Loading struct {}

func NewLoading() *Loading {
	return &Loading{}
}

func (this *Loading) Update() (models.State, error) {
	return models.State_WorldConnecting, nil
}

func (this *Loading) Draw(screen *ebiten.Image) {
}

func (this *Loading) HandleInput() {

}
