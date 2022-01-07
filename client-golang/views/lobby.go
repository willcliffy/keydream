package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/models"
)

type Lobby struct {
	State models.State
}

func NewLobby() *Lobby {
	return &Lobby{
		State: models.State_LobbyConnecting,
	}
}

func (this *Lobby) Update() (models.State, error) {
	return this.State, nil
}

func (this *Lobby) Draw(screen *ebiten.Image) {

}

func (this *Lobby) HandleInput() {

}
