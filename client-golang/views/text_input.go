package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/models"
	"golang.org/x/image/font"
)

type TextInput struct {
	TextBox string
	TextBoxFont font.Face
	X int
	Y int
	W int
	H int
}

func NewTextInput(fonts map[models.FontSize]font.Face, x, y, w, h int) *TextInput {
	return &TextInput{
		TextBoxFont: fonts[models.FontSizeSmall],
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (this *TextInput) Update() error {
	input := inpututil.AppendPressedKeys([]ebiten.Key{})

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		if len(this.TextBox) > 0 {
			this.TextBox = this.TextBox[:len(this.TextBox)-1]
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		this.TextBox += " "
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return nil
	}

	for _, key := range input {
		if inpututil.IsKeyJustPressed(key) &&
			len(this.TextBox) < 10 &&
			len(key.String()) == 1 {
			this.TextBox += key.String()
		}
	}

	return nil
}

func (this *TextInput) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(this.W, this.H)
	img.Fill(color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(this.X), float64(this.Y))
	screen.DrawImage(img, op)

	text.Draw(screen, this.TextBox, this.TextBoxFont, this.X + this.W/5, this.Y + this.H*2/3, color.Black)
}

