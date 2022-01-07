package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/models"
	"github.com/willcliffy/keydream/client/utils"
	"github.com/willcliffy/keydream/client/views"
)

type KeydreamGame struct {
	currentState models.State
	currentView views.View
	views map[models.State]views.View
}

func NewGame() (*KeydreamGame, error) {
	gameFonts, err := models.LoadFonts()
	if err != nil {
		return nil, err
	}

	background, err := views.NewBackground()
	if err != nil {
		return nil, err
	}

	game := &KeydreamGame{
		currentState: models.State_LobbyDisconnected,
	}

	title := views.NewTitleScreen(gameFonts, background)
	lobby := views.NewLobby()
	world := views.NewWorld()

	game.views = map[models.State]views.View{
		models.State_LobbyDisconnected: title,
		models.State_LobbyConnecting:   lobby,
		models.State_LobbyConnected:    lobby,
		models.State_WorldConnecting:   world,
		models.State_WorldConnected:    world,
	}

	game.setState(models.State_LobbyDisconnected)

	return game, nil
}

func (g *KeydreamGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return utils.ScreenWidth, utils.ScreenHeight
}

func (g *KeydreamGame) Update() error {
	state, err := g.currentView.Update()
	if err != nil {
		return err
	}

	if state != g.currentState {
		g.setState(state)
	}

	return nil
}

func (g *KeydreamGame) Draw(screen *ebiten.Image) {
	g.currentView.Draw(screen)
}

func (g *KeydreamGame) setState(state models.State) {
	g.currentState = state
	g.currentView = g.views[state]
	fmt.Printf("State changed to %v\n", state)
}
