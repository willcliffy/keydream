package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/game"
	"github.com/willcliffy/keydream/client/utils"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
