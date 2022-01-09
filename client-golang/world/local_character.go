package world

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/willcliffy/keydream/client/common"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
	world_models "github.com/willcliffy/keydream/client/world/models"
)

type LocalCharacter struct {
	Character

	ID uint8
	Player *common.Player

	NoEquipmentAnimations map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation
	WithSwordAnimations map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation

	// todo - this is hacky. dont be like this
	withSword bool

	Direction objects.CharacterDirection
	State     objects.CharacterState
	Type      objects.CharacterType

	X, Y int64
	LastX, LastY int64
}

func NewCharacter(player *common.Player, charType objects.CharacterType) *LocalCharacter {
	animations := make(map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation)
	for _, state := range objects.CharacterState_values() {
		animations[state] = make(map[objects.CharacterDirection]*views.Animation)
	}

	for _, state := range objects.CharacterState_values() {
		for _, direction := range objects.CharacterDirection_values() {
			frames := make([]*ebiten.Image, 4)
			for i := 1; i <= 4; i++ {
				filePath := fmt.Sprintf("./assets/sprites/rgs_dev/Character without weapon/%s/%s %s%d.png", state.String(), state.String(), direction.String(), i)
				f, err := os.Open(filePath)
				if err != nil {
					panic(err)
				}

				rawImg, _, err := image.Decode(f)
				if err != nil {
					panic(err)
				}

				frames[i-1] = ebiten.NewImageFromImage(rawImg)
			}

			animations[state][direction] = views.NewAnimation(frames, constants.CharacterAnimationSpeed)
		}
	}

	wsAnimations := make(map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation)
	for _, state := range objects.CharacterState_values() {
		wsAnimations[state] = make(map[objects.CharacterDirection]*views.Animation)
	}

	for _, state := range objects.CharacterState_values() {
		for _, direction := range objects.CharacterDirection_values() {
			frames := make([]*ebiten.Image, 4)
			for i := 1; i <= 4; i++ {
				filePath := fmt.Sprintf("./assets/sprites/rgs_dev/Character with sword and shield/%s/%s %s%d.png", state.String(), state.String(), direction.String(), i)
				f, err := os.Open(filePath)
				if err != nil {
					panic(err)
				}

				rawImg, _, err := image.Decode(f)
				if err != nil {
					panic(err)
				}

				frames[i-1] = ebiten.NewImageFromImage(rawImg)
			}

			wsAnimations[state][direction] = views.NewAnimation(frames, constants.CharacterAnimationSpeed)
		}
	}

	return &LocalCharacter{
		Player:    player,

		Direction: objects.CharacterDirection_DOWN,
		State:     objects.CharacterState_IDLE,
		Type:      charType,

		NoEquipmentAnimations: animations,
		WithSwordAnimations: wsAnimations,

		X:     constants.TileSizeScaled,
		Y:     constants.TileSizeScaled,
		LastX: constants.TileSizeScaled,
		LastY: constants.TileSizeScaled,
	}
}

func (c *LocalCharacter) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		c.withSword = !c.withSword
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		c.WalkInDirection(objects.CharacterDirection_LEFT)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		c.WalkInDirection(objects.CharacterDirection_RIGHT)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		c.WalkInDirection(objects.CharacterDirection_UP)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		c.WalkInDirection(objects.CharacterDirection_DOWN)
	} else {
		pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
		if len(pressedKeys) == 0 {
			c.State = objects.CharacterState_IDLE
		} else {
			for _, key := range pressedKeys {
				if c.Direction == objects.KeyToDirection(key) {
					c.WalkInDirection(c.Direction)
					break
				}
			}
		}
	}

	if c.withSword {
		c.WithSwordAnimations[c.State][c.Direction].Update()
	} else {
		c.NoEquipmentAnimations[c.State][c.Direction].Update()
	}
}

func (c *LocalCharacter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)

	var frame *ebiten.Image
	if c.withSword {
		frame = c.WithSwordAnimations[c.State][c.Direction].GetCurrentFrame()
	} else {
		frame = c.NoEquipmentAnimations[c.State][c.Direction].GetCurrentFrame()
	}
	
	screen.DrawImage(frame, op)
}

func (c *LocalCharacter) WalkInDirection(direction objects.CharacterDirection) {
	c.Direction = direction
	c.State = objects.CharacterState_WALK

	switch direction {
	case objects.CharacterDirection_LEFT:
		c.X -= constants.CharacterWalkSpeed
	case objects.CharacterDirection_RIGHT:
		c.X += constants.CharacterWalkSpeed
	case objects.CharacterDirection_UP:
		c.Y -= constants.CharacterWalkSpeed
	case objects.CharacterDirection_DOWN:
		c.Y += constants.CharacterWalkSpeed
	}
}

func (c *LocalCharacter) HasMoved() bool {
	return c.X != c.LastX || c.Y != c.LastY
}

func (c *LocalCharacter) Tick() {
	c.LastX = c.X
	c.LastY = c.Y
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

func (c *LocalCharacter) HandleTock(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleMove(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleAttack(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleJoin(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleLeft(msg *world_models.WorldMessage) {

}

func (c *LocalCharacter) HandleChat(msg *world_models.WorldMessage) {

}
