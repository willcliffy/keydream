package lobby

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/views"
	"golang.org/x/image/font"
)

func NewNameInput(fonts map[models.FontSize]font.Face) *views.TextInput {
	return views.NewTextInput(fonts,
		constants.ScreenWidth/2 - constants.DefaultButtonWidth,
		constants.ScreenHeight/2 - constants.DefaultButtonHeight/2,
		2 * constants.DefaultButtonWidth,
		constants.DefaultButtonHeight)
}

func NewConnectButton(fonts map[models.FontSize]font.Face) *views.Button {
	return views.NewButton("Connect",
		constants.ScreenWidth/2 - constants.DefaultButtonWidth/2,
		constants.ScreenHeight/2 + 2*constants.DefaultButtonHeight,
		constants.DefaultButtonWidth,
		constants.DefaultButtonHeight,
		fonts[models.FontSizeSmall])
}

type Title struct {
	Player *common.Player

	TitleFont font.Face
	Tileset *views.Tileset

	NameInput *views.TextInput
	ConnectButton *views.Button
}

func NewTitleScreen(player *common.Player, fonts map[models.FontSize]font.Face, tileset *views.Tileset) *Title {
	return &Title{
		Player: player,
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
			this.Player.Name = this.NameInput.TextBox
			return models.State_LobbyConnecting, nil
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
			this.ConnectButton.IsMouseOver(ebiten.CursorPosition()) {
			this.Player.Name = this.NameInput.TextBox
			return models.State_LobbyConnecting, nil
		}
	}

	return models.State_LobbyDisconnected, nil
}

func (this *Title) Draw(screen *ebiten.Image) {
	this.Tileset.Draw(screen)
	text.Draw(screen, "Keydream", this.TitleFont, constants.ScreenWidth/3, constants.ScreenHeight/3, color.White)
	this.NameInput.Draw(screen)
	this.ConnectButton.Draw(screen)
}

func (this *Title) HandleInput() error {
	return nil
}
