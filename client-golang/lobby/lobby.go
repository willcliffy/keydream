package lobby

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/views"
	"github.com/willcliffy/keydream/client/common/models"
	"golang.org/x/image/font"
)


func NewBackButton(fonts map[models.FontSize]font.Face) *views.Button {
	return views.NewButton(" Back",
		common.ScreenWidth/2 - common.DefaultButtonWidth/2,
		common.ScreenHeight/2 + 2*common.DefaultButtonHeight,
		common.DefaultButtonWidth,
		common.DefaultButtonHeight,
		fonts[models.FontSizeSmall])
}

type Lobby struct {
	State models.State
	Fonts map[models.FontSize]font.Face
	Tileset *views.Tileset
	BackButton *views.Button
}

func NewLobby(fonts map[models.FontSize]font.Face, tileset *views.Tileset) *Lobby {
	return &Lobby{
		State: models.State_LobbyConnecting,
		Fonts: fonts,
		Tileset: tileset,
		BackButton: NewBackButton(fonts),
	}
}

func (this *Lobby) Update() (models.State, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return models.State_LobbyDisconnected, nil
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		this.BackButton.IsMouseOver(ebiten.CursorPosition()) {
		return models.State_LobbyDisconnected, nil
	}

	if this.State == models.State_LobbyConnecting {
		this.Connect()
	}

	return this.State, nil
}

func (this *Lobby) Draw(screen *ebiten.Image) {
	this.Tileset.Draw(screen)
	this.BackButton.Draw(screen)
}

func (this *Lobby) Connect() {
	// todo - make call to server
}