package views

import "github.com/hajimehoshi/ebiten/v2"

type Title struct {}

func NewTitle() *Title {
	return &Title{}
}

func (this *Title) Update() error {
	return nil
}

func (this *Title) Draw(screen *ebiten.Image) {

}

func (this *Title) HandleInput() {

}
