package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/keydream"
)

func main() {
	game, err := keydream.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(keydream.ScreenWidth, keydream.ScreenHeight)
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}