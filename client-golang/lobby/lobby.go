package lobby

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/models"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
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
	Player *common.Player

	BaseURL string

	State models.State

	Fonts map[models.FontSize]font.Face
	Tileset *views.Tileset
	BackButton *views.Button

	Worlds []WorldView
}

func NewLobby(player *common.Player, baseUrl string, fonts map[models.FontSize]font.Face, tileset *views.Tileset) *Lobby {
	return &Lobby{
		Player: player,
		BaseURL: baseUrl,
		State: models.State_LobbyConnecting,
		Fonts: fonts,
		Tileset: tileset,
		BackButton: NewBackButton(fonts),
	}
}

func (this *Lobby) Update() (models.State, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return models.State_LobbyDisconnected, nil
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if this.BackButton.IsMouseOver(ebiten.CursorPosition()) {
			return models.State_LobbyDisconnected, nil
		}
		for _, w := range this.Worlds {
			if w.IsMouseOver(ebiten.CursorPosition()) {
				fmt.Printf("joining world %d\n", w.Data.ID)
				return models.State_WorldConnecting, nil
			}
		}
	}

	if this.State == models.State_LobbyConnecting {
		if err := this.Connect(); err != nil {
			return models.State_LobbyDisconnected, err
		}

		this.State = models.State_LobbyConnected
	} else if this.State == models.State_LobbyConnected {
		for _, w := range this.Worlds {
			if err := w.Update(); err != nil {
				return models.State_LobbyDisconnected, nil
			}
		}
	}

	return this.State, nil
}

func (this *Lobby) Draw(screen *ebiten.Image) {
	this.Tileset.Draw(screen)

	this.BackButton.Draw(screen)
	for _, w := range this.Worlds {
		w.Draw(screen)
	}

	if this.State == models.State_LobbyConnected {
		text.Draw(screen, fmt.Sprintf("Connected to lobby as: %s", this.Player.Name), this.Fonts[models.FontSizeTiny], common.TileSize, common.TileSize, color.Black)
	}
}

func (this *Lobby) Connect() error {
	payload := objects.ConnectRequest{
		Name: "",
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// todo - make call to server
	res, err := http.DefaultClient.Post(
		this.BaseURL + "/api/v1/connect",
		"application/json",
		bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if res.StatusCode < 200 && res.StatusCode >= 300 {
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error from lobby, couldnt unmarshal response: %s", err)
		}
		return fmt.Errorf("got bad response from lobby %d: %s", res.StatusCode, string(resBody))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response objects.ConnectResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	y := common.ScreenHeight / 8
	for _, w := range response.Worlds {
		this.Worlds = append(this.Worlds, *NewWorldView(&w, this.Fonts[models.FontSizeSmall], y))
		y += common.DefaultButtonHeight
	}

	fmt.Printf("got response from lobby: %+v\n", response)

	return nil
}
