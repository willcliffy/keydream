package world

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
	world_models "github.com/willcliffy/keydream/client/world/models"
)

type RemoteCharacter struct {
	Character

	Name string
	ID uint8

	NoEquipmentAnimations map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation
	WithSwordAnimations map[objects.CharacterState]map[objects.CharacterDirection]*views.Animation

	// todo - this is hacky. dont be like this
	withSword bool

	Direction objects.CharacterDirection
	State     objects.CharacterState
	Type      objects.CharacterType

	X, Y int64

	Targets []objects.Position
}

func NewRemoteCharacter(id uint8, characterName string, x, y int64) *RemoteCharacter {
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

	return &RemoteCharacter{
		Name: characterName,
		ID: id,

		X: x,
		Y: y,

		NoEquipmentAnimations: animations,
		WithSwordAnimations: animations,

		Direction: objects.CharacterDirection_DOWN,
		State:     objects.CharacterState_IDLE,
		Type:      objects.CharacterType_REMOTE,
	}
}

func (r *RemoteCharacter) Update() {
	if len(r.Targets) == 0 {
		r.State = objects.CharacterState_IDLE
		r.NoEquipmentAnimations[r.State][r.Direction].Update()
	} else {
		if r.X - r.Targets[0].X > constants.CharacterWalkSpeed {
			r.X -= constants.CharacterWalkSpeed
		} else if r.X - r.Targets[0].X < -constants.CharacterWalkSpeed {
			r.X += constants.CharacterWalkSpeed
		} else {
			r.X = r.Targets[0].X
		}

		if r.Y - r.Targets[0].Y > constants.CharacterWalkSpeed {
			r.Y -= constants.CharacterWalkSpeed
		} else if r.Y - r.Targets[0].Y < -constants.CharacterWalkSpeed {
			r.Y += constants.CharacterWalkSpeed
		} else {
			r.Y = r.Targets[0].Y
		}
	}
}

func (r *RemoteCharacter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)

	var frame *ebiten.Image
	if r.withSword {
		frame = r.WithSwordAnimations[r.State][r.Direction].GetCurrentFrame()
	} else {
		frame = r.NoEquipmentAnimations[r.State][r.Direction].GetCurrentFrame()
	}
	
	screen.DrawImage(frame, op)
}


func (r *RemoteCharacter) HandleMessage(msg *world_models.WorldMessage) error {
	switch msg.Type {
	case world_models.WorldMessageType_MOVE:
		if len(msg.Params) != 2 {
			return fmt.Errorf("expected 2 params, got %+v", msg.Params)
		}

		x, err := strconv.ParseInt(msg.Params[0], 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse x param: %s", err)
		}

		y, err := strconv.ParseInt(msg.Params[1], 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse y param: %s", err)
		}

		r.handleMove(x, y)
		return nil
	default:
		return fmt.Errorf("unknown message type for remote character: %s", msg.Type)
	}
}

func (r *RemoteCharacter) handleMove(x, y int64) {
	r.Targets = append(r.Targets, objects.Position{X: x, Y: y})
}

