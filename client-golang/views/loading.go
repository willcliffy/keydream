package views

import "github.com/hajimehoshi/ebiten/v2"

type Loading struct {}

func NewLoading() *Loading {
	return &Loading{}
}

func (this *Loading) Update() error {
	return nil
}

func (this *Loading) Draw(screen *ebiten.Image) {

}

func (this *Loading) HandleInput() {

}
