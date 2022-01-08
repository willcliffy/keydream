package world

import (
	"fmt"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
)

type World struct {
	State models.State
	BaseURL string
	Character *Character
	Background *views.Tileset
	
	conn *net.UDPConn
}

func NewWorld(player *common.Player, baseUrl string, tileset *views.Tileset) *World {
	return &World{
		State: models.State_WorldConnecting,
		BaseURL: baseUrl,
		Character: NewCharacter(player, objects.CharacterType_Local),
		Background: tileset,
	}
}

func (this *World) Update() (models.State, error) {
	if this.State == models.State_WorldConnecting {
		if err := this.Connect(); err != nil {
			return models.State_LobbyConnected, err
		} else {
			return models.State_WorldConnected, nil
		}
	} else if this.State == models.State_WorldConnected {
		this.Character.Update()
		return models.State_WorldConnected, nil
	}

	return this.State, nil
}

func (this *World) Draw(screen *ebiten.Image) {
	this.Background.Draw(screen)
	this.Character.Draw(screen)
}

func (this *World) HandleInput() error {
	return nil
}

func (this *World) Connect() error {
	ServerAddr, err := net.ResolveUDPAddr("udp", this.BaseURL)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		return err
	}
	this.conn = conn

	_, err = this.conn.Write([]byte("join " + this.Character.Player.Name + " \n"))
	if err != nil {
		fmt.Printf("error writing to gameserver: %s", err)
	}

	buf := make([]byte, 128)
	_, _, err = this.conn.ReadFrom(buf)
	if err != nil {
		fmt.Printf("error reading from gameserver: %s\n", err)
	}

	fmt.Printf("%s", buf)

	this.State = models.State_WorldConnected
	return nil
}

func (this *World) Dispose() error {
	this.conn.Close()
	this.State = models.State_LobbyConnected
	return nil
}
