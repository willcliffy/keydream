package keydream

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type KeydreamGame struct {
	
}

func NewGame() (*KeydreamGame, error) {
	game := &KeydreamGame{}
	return game, nil
}

func (g *KeydreamGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *KeydreamGame) Update() error {
	return nil
}

func (g *KeydreamGame) Draw(screen *ebiten.Image) {
	
}
