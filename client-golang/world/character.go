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
)

type Character struct {
	Player *common.Player

	Animations map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation

	Direction objects.CharacterDirection
	State     objects.CharacterState
	Type      objects.CharacterType

	X, Y int64
	LastX, LastY int64
}

func NewCharacter(player *common.Player, charType objects.CharacterType) *Character {
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

	return &Character{
		Player:    player,

		Direction: objects.CharacterDirection_Down,
		State:     objects.CharacterState_Idle,
		Type:      charType,

		Animations: animations,

		X:     constants.TileSizeScaled,
		Y:     constants.TileSizeScaled,
		LastX: constants.TileSizeScaled,
		LastY: constants.TileSizeScaled,
	}
}

func (c *Character) Update() {
	c.Animations[c.State][c.Direction].Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		c.WalkInDirection(objects.CharacterDirection_Left)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		c.WalkInDirection(objects.CharacterDirection_Right)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		c.WalkInDirection(objects.CharacterDirection_Up)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		c.WalkInDirection(objects.CharacterDirection_Down)
	} else {
		pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
		if len(pressedKeys) == 0 {
			c.State = objects.CharacterState_Idle
			return
		}

		for _, key := range pressedKeys {
			if c.Direction == objects.CharacterDirection(key) {
				c.WalkInDirection(c.Direction)
				break
			}
		}
	}
}

func (c *Character) Draw(screen *ebiten.Image) {
	animation := c.Animations[c.State][c.Direction]
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)

	screen.DrawImage(animation.GetCurrentFrame(), op)
}

func (c *Character) WalkInDirection(direction objects.CharacterDirection) {
	c.Direction = direction
	c.State = objects.CharacterState_Walk

	switch direction {
	case objects.CharacterDirection_Left:
		c.X -= constants.LocalCharacterWalkSpeed
	case objects.CharacterDirection_Right:
		c.X += constants.LocalCharacterWalkSpeed
	case objects.CharacterDirection_Up:
		c.Y -= constants.LocalCharacterWalkSpeed
	case objects.CharacterDirection_Down:
		c.Y += constants.LocalCharacterWalkSpeed
	}
}
