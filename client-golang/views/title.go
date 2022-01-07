package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/models"
	"github.com/willcliffy/keydream/client/utils"
	"golang.org/x/image/font"
)

func NewNameInput(fonts map[models.FontSize]font.Face) *TextInput {
	return NewTextInput(fonts,
		utils.ScreenWidth/2 - utils.DefaultButtonWidth,
		utils.ScreenHeight/2 - utils.DefaultButtonHeight/2,
		2 * utils.DefaultButtonWidth,
		utils.DefaultButtonHeight)
}

func NewConnectButton(fonts map[models.FontSize]font.Face) *Button {
	return NewButton("Connect",
		utils.ScreenWidth/2 - utils.DefaultButtonWidth/2,
		utils.ScreenHeight/2 + 2*utils.DefaultButtonHeight,
		utils.DefaultButtonWidth,
		utils.DefaultButtonHeight,
		fonts[models.FontSizeSmall])
}

type Title struct {
	TitleFont font.Face
	Background *Background

	NameInput *TextInput
	ConnectButton *Button
}

func NewTitleScreen(fonts map[models.FontSize]font.Face, background *Background) *Title {
	return &Title{
		TitleFont: fonts[models.FontSizeLarge],
		Background: background,
		ConnectButton: NewConnectButton(fonts),
		NameInput: NewNameInput(fonts),
	}
}

func (this *Title) Update() (models.State, error) {
	if err := this.NameInput.Update(); err != nil {
		return models.State_LobbyDisconnected, err
	}

	if len(this.NameInput.TextBox) > 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			return models.State_LobbyConnecting, nil
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
			this.ConnectButton.IsMouseOver(ebiten.CursorPosition()) {
			return models.State_LobbyConnecting, nil
		}
	}

	return models.State_LobbyDisconnected, nil
}

func (this *Title) Draw(screen *ebiten.Image) {
	this.Background.Draw(screen)
	text.Draw(screen, "Keydream", this.TitleFont, utils.ScreenWidth/3, utils.ScreenHeight/3, color.White)
	this.NameInput.Draw(screen)
	this.ConnectButton.Draw(screen)
}
