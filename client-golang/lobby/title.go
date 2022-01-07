package lobby

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/views"
	"golang.org/x/image/font"
)

func NewNameInput(fonts map[models.FontSize]font.Face) *views.TextInput {
	return views.NewTextInput(fonts,
		common.ScreenWidth/2 - common.DefaultButtonWidth,
		common.ScreenHeight/2 - common.DefaultButtonHeight/2,
		2 * common.DefaultButtonWidth,
		common.DefaultButtonHeight)
}

func NewConnectButton(fonts map[models.FontSize]font.Face) *views.Button {
	return views.NewButton("Connect",
		common.ScreenWidth/2 - common.DefaultButtonWidth/2,
		common.ScreenHeight/2 + 2*common.DefaultButtonHeight,
		common.DefaultButtonWidth,
		common.DefaultButtonHeight,
		fonts[models.FontSizeSmall])
}

type Title struct {
	TitleFont font.Face
	Tileset *views.Tileset

	NameInput *views.TextInput
	ConnectButton *views.Button
}

func NewTitleScreen(fonts map[models.FontSize]font.Face, tileset *views.Tileset) *Title {
	return &Title{
		TitleFont: fonts[models.FontSizeLarge],
		Tileset: tileset,
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
	this.Tileset.Draw(screen)
	text.Draw(screen, "Keydream", this.TitleFont, common.ScreenWidth/3, common.ScreenHeight/3, color.White)
	this.NameInput.Draw(screen)
	this.ConnectButton.Draw(screen)
}
