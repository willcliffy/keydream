package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/views"
)

// todo - this shouldnt be hardcoded eventually
const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type View interface {
	//Layout(outsideWidth, outsideHeight int) (int, int)
	Update() error
	Draw(screen *ebiten.Image)
}

type KeydreamGame struct {
	currentView View
	views map[State]View
}

func NewGame() (*KeydreamGame, error) {
	game := &KeydreamGame{
		views: map[State]View{
			State_Disconnected:    views.NewTitle(),
			State_LobbyConnected:  views.NewLobby(),
			State_WorldConnecting: views.NewLoading(),
			State_WorldConnected:  views.NewWorld(),
		},
	}

	game.currentView = game.views[State_Disconnected]

	return game, nil
}

func (g *KeydreamGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *KeydreamGame) Update() error {
	return g.currentView.Update()
}

func (g *KeydreamGame) Draw(screen *ebiten.Image) {
	g.currentView.Draw(screen)
}
