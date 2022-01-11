package world

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
)

type LocalCharacter struct {
	Character

	ID uint8
	Player *common.Player
	Type      objects.CharacterType

	Animation *views.CharacterAnimation

	x, y float64
	lastX, lastY float64
}

func NewCharacter(player *common.Player, charType objects.CharacterType) *LocalCharacter {
	animation := views.NewCharacterAnimation()

	return &LocalCharacter{
		Player:    player,
		Type:      charType,

		Animation: &animation,

		x:         constants.TileSizeScaled,
		y:         constants.TileSizeScaled,
		lastX:     constants.TileSizeScaled,
		lastY:     constants.TileSizeScaled,
	}
}

func (c *LocalCharacter) Update() {
	if directions := c.directionsPressed(); len(directions) == 0 {
		c.Animation.Idle()
	} else {
		for _, direction := range directions {
			c.moveInDirection(direction)
			c.Animation.Walk(direction)
		}
	}

	c.Animation.Update()
}

func (c *LocalCharacter) directionsPressed() []objects.CharacterDirection {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	directions := make([]objects.CharacterDirection, 0)
	for _, key := range pressedKeys {
		dir := objects.KeyToDirection(key)
		if dir != objects.CharacterDirection_NONE && !dir.ContainedIn(directions) && !dir.OppositeContainedIn(directions) {
			directions = append(directions, dir)
		}
	}

	return directions
}

func (c *LocalCharacter) moveInDirection(direction objects.CharacterDirection) {
	switch direction {
	case objects.CharacterDirection_LEFT:
		c.x -= constants.CharacterWalkSpeed
	case objects.CharacterDirection_RIGHT:
		c.x += constants.CharacterWalkSpeed
	case objects.CharacterDirection_UP:
		c.y -= constants.CharacterWalkSpeed
	case objects.CharacterDirection_DOWN:
		c.y += constants.CharacterWalkSpeed
	}
}

func (c *LocalCharacter) Draw(screen *ebiten.Image) {
	c.Animation.Draw(screen, c.x, c.y)
}

func (c *LocalCharacter) HasMoved() bool {
	return c.x != c.lastX || c.y != c.lastY
}

func (c *LocalCharacter) Tick() {
	c.lastX = c.x
	c.lastY = c.y
}

func (c *LocalCharacter) HandleMessage(msg *objects.WorldMessage) error {
	switch msg.Type {
	case objects.WorldMessageType_JOIN:
		c.HandleJoin(msg)
	case objects.WorldMessageType_MOVE:
		c.HandleMove(msg)
	case objects.WorldMessageType_LEFT:
		c.HandleLeft(msg)
	default:
		return fmt.Errorf("Unknown message type for local player: %s", msg.Type)
	}

	return nil
}

func (c *LocalCharacter) HandleMove(msg *objects.WorldMessage) {

}

func (c *LocalCharacter) HandleJoin(msg *objects.WorldMessage) {

}

func (c *LocalCharacter) HandleLeft(msg *objects.WorldMessage) {

}
