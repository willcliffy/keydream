package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/models"
)

type Lobby struct {}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (this *Lobby) Update() (models.State, error) {
	return models.State_LobbyConnected, nil
}

func (this *Lobby) Draw(screen *ebiten.Image) {

}

func (this *Lobby) HandleInput() {

}
