package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/views"
	"github.com/willcliffy/keydream/client/lobby"
	"github.com/willcliffy/keydream/client/world"
)

type KeydreamGame struct {
	currentState models.State
	currentView common.View
	views map[models.State]common.View
}

func NewGame() (*KeydreamGame, error) {
	gameFonts, err := models.LoadFonts()
	if err != nil {
		return nil, err
	}

	tileset, err := views.NewTileset()
	if err != nil {
		return nil, err
	}

	game := &KeydreamGame{
		currentState: models.State_LobbyDisconnected,
	}

	title := lobby.NewTitleScreen(gameFonts, tileset)
	lobby := lobby.NewLobby(gameFonts, tileset)
	world := world.NewWorld()

	game.views = map[models.State]common.View{
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
	return common.ScreenWidth, common.ScreenHeight
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
