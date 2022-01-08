package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	ButtonFont font.Face
	Text string
	X int
	Y int
	W int
	H int
}

func NewButton(text string, x, y, w, h int, buttonFont font.Face) *Button {
	return &Button{
		ButtonFont: buttonFont,
		Text: text,
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (this *Button) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(this.W, this.H)
	img.Fill(color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(this.X), float64(this.Y))
	screen.DrawImage(img, op)

	// TODO: Fix this. This... it's horrid.
	text.Draw(screen, this.Text, this.ButtonFont, this.X + this.W/5, this.Y + this.H*2/3, color.Black)
}

func (this *Button) IsMouseOver(x, y int) bool {
	return x >= this.X && x <= this.X + this.W && y >= this.Y && y <= this.Y + this.H
}
