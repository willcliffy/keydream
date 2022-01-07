package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/game"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(game.Layout(0, 0))
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
