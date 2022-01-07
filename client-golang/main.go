package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/game"
	"github.com/willcliffy/keydream/client/common"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Keydream")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
