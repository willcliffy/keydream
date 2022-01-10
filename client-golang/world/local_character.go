package world

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
	world_models "github.com/willcliffy/keydream/client/world/models"
	world_utils "github.com/willcliffy/keydream/client/world/utils"
)

type LocalCharacter struct {
	Character

	ID uint8
	Player *common.Player

	IdleAnimation *views.Animation
	WalkAnimation *views.Animation

	Direction objects.CharacterDirection
	State     objects.CharacterState
	Type      objects.CharacterType

	x, y float64
	lastX, lastY float64
}

func NewCharacter(player *common.Player, charType objects.CharacterType) *LocalCharacter {
	return &LocalCharacter{
		Player:    player,

		Direction: objects.CharacterDirection_DOWN,
		State:     objects.CharacterState_IDLE,
		Type:      charType,

		IdleAnimation: world_utils.LoadIdleAnimations(),
		WalkAnimation: world_utils.LoadWalkAnimations(),

		x:     constants.TileSizeScaled,
		y:     constants.TileSizeScaled,
		lastX: constants.TileSizeScaled,
		lastY: constants.TileSizeScaled,
	}
}

func (c *LocalCharacter) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		c.walkInDirection(objects.CharacterDirection_LEFT)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		c.walkInDirection(objects.CharacterDirection_RIGHT)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		c.walkInDirection(objects.CharacterDirection_UP)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		c.walkInDirection(objects.CharacterDirection_DOWN)
	} else {
		pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
		if len(pressedKeys) == 0 {
			c.State = objects.CharacterState_IDLE
			c.IdleAnimation.Update(c.Direction)
		} else {
			for _, key := range pressedKeys {
				if c.Direction == objects.KeyToDirection(key) {
					c.walkInDirection(c.Direction)
					break
				}
			}
		}
	}
}

func (c *LocalCharacter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.x, c.y)
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)
	
	switch c.State {
	case objects.CharacterState_IDLE:
		screen.DrawImage(c.IdleAnimation.GetCurrentFrame(c.Direction), op)
	case objects.CharacterState_WALK:
		screen.DrawImage(c.WalkAnimation.GetCurrentFrame(c.Direction), op)
	}
}

func (c *LocalCharacter) walkInDirection(direction objects.CharacterDirection) {
	c.Direction = direction
	c.State = objects.CharacterState_WALK

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

	c.WalkAnimation.Update(direction)
}

func (c *LocalCharacter) HasMoved() bool {
	return c.x != c.lastX || c.y != c.lastY
}

func (c *LocalCharacter) Tick() {
	c.lastX = c.x
	c.lastY = c.y
}

func (c *LocalCharacter) HandleMessage(msg *world_models.WorldMessage) error {
	switch msg.Type {
	case world_models.WorldMessageType_JOIN:
		c.HandleJoin(msg)
	case world_models.WorldMessageType_MOVE:
		c.HandleMove(msg)
	case world_models.WorldMessageType_LEFT:
		c.HandleLeft(msg)
	default:
		return fmt.Errorf("Unknown message type for local player: %s", msg.Type)
	}

	return nil
}

func (c *LocalCharacter) HandleMove(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleJoin(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleLeft(msg *world_models.WorldMessage) {

}
