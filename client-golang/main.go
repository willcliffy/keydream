package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/game"
	"github.com/willcliffy/keydream/client/models"
	"github.com/willcliffy/keydream/client/utils"
	"github.com/willcliffy/keydream/client/views"
)

func main() {
	gameFonts, err := game.LoadFonts()
	if err != nil {
		log.Fatal(err)
	}

	views := map[models.State]views.View{
		models.State_Disconnected:    views.NewTitle(gameFonts[game.FontSizeLarge]),
		models.State_LobbyConnected:  views.NewLobby(),
		models.State_WorldConnecting: views.NewLoading(),
		models.State_WorldConnected:  views.NewWorld(),
	}

	game, err := game.NewGame(views)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
