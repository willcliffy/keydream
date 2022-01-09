package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
)

type WorldView struct {
	State models.State
	BaseURL string
	Character *LocalCharacter
	Background *views.Tileset
	
	transport *WorldTransport
}

func NewWorld(player *common.Player, baseUrl string, tileset *views.Tileset) *WorldView {
	character := NewCharacter(player, objects.CharacterType_LOCAL)
	transport := NewWorldTransport(character, baseUrl)

	return &WorldView{
		State: models.State_WorldConnecting,
		BaseURL: baseUrl,
		Character: character,
		Background: tileset,

		transport: transport,
	}
}

func (this *WorldView) Update() (models.State, error) {
	if this.State == models.State_WorldConnecting {
		if err := this.transport.Connect(); err != nil {
			return models.State_LobbyConnected, err
		} else {
			this.State = models.State_WorldConnected
		}
	} else if this.State == models.State_WorldConnected {
		this.Character.Update()
		if err := this.transport.Update(); err != nil {
			return models.State_LobbyConnected, err
		}
	}

	return this.State, nil
}

func (this *WorldView) Draw(screen *ebiten.Image) {
	this.Background.Draw(screen)
	this.Character.Draw(screen)
	for _, rc := range this.transport.RemoteCharacters {
		rc.Draw(screen)
	}
}

func (this *WorldView) HandleInput() error {
	return nil
}

