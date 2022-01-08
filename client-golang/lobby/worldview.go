package lobby

import (
	"fmt"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
)

type WorldView struct {
	Data *objects.WorldBroadcast

	Font font.Face
	X int
	Y int
	W int
	H int
}

func NewWorldView(data *objects.WorldBroadcast, font font.Face, y int) *WorldView {
	return &WorldView{
		Data: data,
		Font: font,
		X: constants.ScreenWidth / 2 - 1.5 * constants.DefaultButtonWidth,
		Y: y,
		W: 3 * constants.DefaultButtonWidth,
		H: constants.DefaultButtonHeight,
	}
}

func (this *WorldView) Update() error {
	return nil
}

func (this *WorldView) Draw(screen *ebiten.Image) {
	background := ebiten.NewImage(this.W, this.H)
	background.Fill(color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(this.X), float64(this.Y))

	screen.DrawImage(background, op)
	text.Draw(screen, fmt.Sprintf("World %d", this.Data.ID), this.Font,
		this.X + this.W/8,
		this.Y + this.H*3/5,
		color.Black)
	text.Draw(screen, fmt.Sprintf("%d / 20 players", this.Data.NumPlayers), this.Font,
		this.X + this.W/2,
		this.Y + this.H*3/5,
		color.Black)
}

func (this *WorldView) IsMouseOver(x, y int) bool {
	return x >= this.X && x <= this.X + this.W && y >= this.Y && y <= this.Y + this.H
}
