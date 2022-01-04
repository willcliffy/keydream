package views

import "github.com/hajimehoshi/ebiten/v2"

type World struct {}

func NewWorld() *World {
	return &World{}
}

func (this *World) Update() error {
	return nil
}

func (this *World) Draw(screen *ebiten.Image) {

}

func (this *World) HandleInput() {

}
