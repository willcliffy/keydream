package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/game"
)

// todo - this shouldnt be hardcoded eventually
const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
