package objects

import (
	"fmt"
	"strconv"
	"strings"
)

type WorldMessageType string
const (
	WorldMessageType_JOIN = "join"
	WorldMessageType_LEFT = "left"
	WorldMessageType_MOVE = "move"
	WorldMessageType_TOCK = "tock"
)

type WorldMessage struct {
	Type string
	CharacterID uint8
	Params []string
}

func NewWorldMessage(rawMsg string) *WorldMessage {
	msg := strings.Split(strings.Split(rawMsg, "\n")[0], " ")

	charID, err := strconv.ParseUint(msg[1], 10, 8)
	if err != nil {
		fmt.Printf("bad world message: %s", rawMsg)
		return nil
	}

	return &WorldMessage{
		Type: msg[0],
		CharacterID: uint8(charID),
		Params: msg[2:],
	}
}

func (this WorldMessage) String() string {
	return fmt.Sprintf("%s %d %s", this.Type, this.CharacterID, strings.Join(this.Params, " "))
}

