package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/models"
	"github.com/willcliffy/keydream/client/utils"
	"golang.org/x/image/font"
)

type Title struct {
	TitleFont font.Face
	Background *Background

	TextBox string
	ConnectButton string
}

func NewTitle(titleFont font.Face, background *Background) *Title {
	return &Title{
		TitleFont: titleFont,
		Background: background,
	}
}

func (this *Title) Update() (models.State, error) {
	return models.State_Disconnected, nil
}

func (this *Title) Draw(screen *ebiten.Image) {
	this.Background.Draw(screen)
	text.Draw(screen, "Keydream", this.TitleFont, utils.ScreenWidth/3, utils.ScreenHeight/3, color.White)
}

func (this *Title) HandleInput() {

}
