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

func NewGame(views map[models.State]views.View) (*KeydreamGame, error) {
	game := &KeydreamGame{
		currentState: models.State_Disconnected,
		views: views,
	}

	game.setState(models.State_Disconnected)

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
