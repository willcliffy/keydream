package world

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/world/models"
)

type WorldTransport struct {
	baseURL string
	conn *net.UDPConn

	lastTick time.Time

	localCharacter *LocalCharacter
	RemoteCharacters map[uint8]*RemoteCharacter
}

func NewWorldTransport(character *LocalCharacter, baseUrl string) *WorldTransport {
	return &WorldTransport{
		baseURL: baseUrl,
		localCharacter: character,
		RemoteCharacters: make(map[uint8]*RemoteCharacter),
	}
}

func (this *WorldTransport) Connect() error {
	ServerAddr, err := net.ResolveUDPAddr("udp", this.baseURL)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		return err
	}
	this.conn = conn

	_, err = this.conn.Write([]byte("join " + this.localCharacter.Player.Name + " \n"))
	if err != nil {
		fmt.Printf("error writing to gameserver: %s", err)
	}

	buf := make([]byte, 128)
	_, _, err = this.conn.ReadFrom(buf) // todo - this could hang forever
	if err != nil {
		fmt.Printf("error reading from gameserver: %s\n", err)
	}

	msg := strings.Split(string(buf), " ")
	pID, err := strconv.ParseUint(msg[0], 10, 8)
	if err != nil {
		return err
	}
	this.localCharacter.ID = uint8(pID)

	fmt.Printf("Connected to world with player ID: %d\n", this.localCharacter.ID)

	go func() {
		buf := make([]byte, 128)
		for {
			n, _, err := this.conn.ReadFrom(buf)
			if err != nil {
				// todo - check if the connection is closed. if so, just return
				fmt.Printf("error reading from gameserver: %s\n", err)
				return
			}

			if n == 0 {
				break
			}

			this.HandleMessage(string(buf))
		}
	}()

	return nil
}

func (this *WorldTransport) Update() error {
	if time.Since(this.lastTick) > constants.WorldTickRate {
		this.lastTick = time.Now()
		if err := this.Tick(); err != nil {
			return err
		}
	}

	for _, remoteCharacter := range this.RemoteCharacters {
		remoteCharacter.Update()
	}

	return nil
}

func (this *WorldTransport) Tick() error {
	if this.localCharacter.HasMoved() {
		_, err := this.conn.Write([]byte(fmt.Sprintf("move %d %d %d \n", this.localCharacter.ID, this.localCharacter.X, this.localCharacter.Y)))
		if err != nil {
			return err
		}
	} else {
		_, err := this.conn.Write([]byte(fmt.Sprintf("tick %d \n", this.localCharacter.ID)))
		if err != nil {
			return err
		}
	}

	this.localCharacter.Tick()

	return nil
}

func (this *WorldTransport) HandleMessage(msg string) {
	if strings.HasPrefix(msg, "tock") {
		return
	}

	worldMsg := world_models.NewWorldMessage(msg)

	if worldMsg == nil {
		fmt.Printf("bad world message: %s", msg)
	} else if worldMsg.CharacterID == this.localCharacter.ID {
		if err := this.localCharacter.HandleMessage(worldMsg); err != nil {
			fmt.Printf("error handling local character message: %s", err)
		}
	} else {
		switch worldMsg.Type {
		case world_models.WorldMessageType_JOIN:
			if len(worldMsg.Params) != 3 {
				fmt.Printf("ignored message - bad join message: %s", msg)
				return
			}

			name := worldMsg.Params[0]

			x, err := strconv.ParseInt(worldMsg.Params[1], 10, 64)
			if err != nil {
				fmt.Printf("ignored message - bad x in join message: '%s' in '%s'", worldMsg.Params[1], msg)
				return
			}

			y, err := strconv.ParseInt(worldMsg.Params[2], 10, 64)
			if err != nil {
				fmt.Printf("ignored message - bad y in join message: '%s' in '%s'", worldMsg.Params[2], msg)
				return
			}

			this.RemoteCharacters[worldMsg.CharacterID] = NewRemoteCharacter(worldMsg.CharacterID, name, x, y)
		case world_models.WorldMessageType_MOVE:
			rc, ok := this.RemoteCharacters[worldMsg.CharacterID]
			if !ok {
				fmt.Printf("ignored message - received move for unknown character: %d\n", worldMsg.CharacterID)
				return
			}

			if err := rc.HandleMessage(worldMsg); err != nil {
				fmt.Printf("error handling move message: %s\n", err)
			}
		case world_models.WorldMessageType_LEFT:
			this.RemoteCharacters[worldMsg.CharacterID] = nil
		default:
			fmt.Printf("ignored message - unknown message type for remote player: %sn", worldMsg.Type)
		}
	}
}

func (this *WorldTransport) Close() {
	this.conn.Close()
}
