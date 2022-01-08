package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/views"
	"github.com/willcliffy/keydream/client/lobby"
	"github.com/willcliffy/keydream/client/world"
)

type KeydreamGame struct {
	currentState models.State
	currentView common.KeydreamView
	views map[models.State]common.KeydreamView
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

	player := &common.Player{}

	title := lobby.NewTitleScreen(player, gameFonts, tileset)
	lobby := lobby.NewLobby(player, "http://lobby.keydream.tk", gameFonts, tileset)
	world := world.NewWorld(player, "world.keydream.tk:80", gameFonts, tileset)

	game.views = map[models.State]common.KeydreamView{
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
	return constants.ScreenWidth, constants.ScreenHeight
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
