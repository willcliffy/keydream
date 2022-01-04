package views

import "github.com/hajimehoshi/ebiten/v2"

type Lobby struct {}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (this *Lobby) Update() error {
	return nil
}

func (this *Lobby) Draw(screen *ebiten.Image) {

}

func (this *Lobby) HandleInput() {

}
