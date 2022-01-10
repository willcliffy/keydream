package world

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
	world_models "github.com/willcliffy/keydream/client/world/models"
	world_utils "github.com/willcliffy/keydream/client/world/utils"
)

type RemoteCharacter struct {
	Character

	Name string
	ID uint8

	IdleAnimation *views.Animation
	WalkAnimation *views.Animation

	characterDirection objects.CharacterDirection
	characterState     objects.CharacterState
	characterType      objects.CharacterType

	x, y float64
	speedX, speedY float64

	targets []objects.Position
}

func NewRemoteCharacter(id uint8, characterName string, x, y float64) *RemoteCharacter {
	return &RemoteCharacter{
		Name: characterName,
		ID: id,
		x: x,
		y: y,
		characterDirection: objects.CharacterDirection_DOWN,
		characterState: objects.CharacterState_IDLE,
		characterType: objects.CharacterType_REMOTE,
		speedX: constants.RemoteCharacterMinWalkSpeed - constants.RemoteCharacterWalkAcceleration,
		speedY: constants.RemoteCharacterMinWalkSpeed - constants.RemoteCharacterWalkAcceleration,
		targets: []objects.Position{},
		IdleAnimation: world_utils.LoadIdleAnimations(),
		WalkAnimation: world_utils.LoadWalkAnimations(),
	}
}

func (r *RemoteCharacter) Update() {
	if len(r.targets) != 0 {
		r.walkTowardsTarget()
	} else {
		r.IdleAnimation.Update(r.characterDirection)
	}
}

func (r *RemoteCharacter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.x, r.y)
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)
	
	switch r.characterState {
	case objects.CharacterState_IDLE:
		screen.DrawImage(r.IdleAnimation.GetCurrentFrame(r.characterDirection), op)
	case objects.CharacterState_WALK:
		screen.DrawImage(r.WalkAnimation.GetCurrentFrame(r.characterDirection), op)
	}
}

func (r *RemoteCharacter) HandleMessage(msg *world_models.WorldMessage) error {
	switch msg.Type {
	case world_models.WorldMessageType_MOVE:
		if len(msg.Params) != 2 {
			return fmt.Errorf("expected 2 params, got %+v", msg.Params)
		}

		x, err := strconv.ParseFloat(msg.Params[0], 64)
		if err != nil {
			return fmt.Errorf("could not parse x param: %s", err)
		}

		y, err := strconv.ParseFloat(msg.Params[1], 64)
		if err != nil {
			return fmt.Errorf("could not parse y param: %s", err)
		}

		r.targets = append(r.targets, objects.Position{X: x, Y: y})
		return nil
	default:
		return fmt.Errorf("unknown message type for remote character: %s", msg.Type)
	}
}

func (r *RemoteCharacter) walkTowardsTarget() {
	targetX := r.targets[0].X
	targetY := r.targets[0].Y

	if targetX - r.x < -constants.RemoteCharacterAlpha {
		r.speedX = -constants.RemoteCharacterMaxWalkSpeed
	} else if targetX - r.x > constants.RemoteCharacterAlpha {
		r.speedX = constants.RemoteCharacterMaxWalkSpeed
	} else {
		r.x = targetX
		r.speedX = 0
	}

	if targetY - r.y < -constants.RemoteCharacterAlpha {
		r.speedY = -constants.RemoteCharacterMaxWalkSpeed
	} else if targetY - r.y > constants.RemoteCharacterAlpha {
		r.speedY = constants.RemoteCharacterMaxWalkSpeed
	} else {
		r.y = targetY
		r.speedY = 0
	}

	if r.x == targetX && r.y == targetY {
		r.targets = r.targets[1:]
		if len(r.targets) == 0 {
			r.characterState = objects.CharacterState_IDLE
			r.speedX = 0
			r.speedY = 0
		}
		return
	}

	r.x += r.speedX
	r.y += r.speedY

	if r.speedX < constants.RemoteCharacterMaxWalkSpeed {
		r.speedX += constants.RemoteCharacterWalkAcceleration
	}
	if r.speedY < constants.RemoteCharacterMaxWalkSpeed {
		r.speedY += constants.RemoteCharacterWalkAcceleration
	}

	r.characterState = objects.CharacterState_WALK
	if math.Abs(r.speedX) > math.Abs(r.speedY) {
		if r.x > targetX {
			r.characterDirection = objects.CharacterDirection_LEFT
		} else if r.x < targetX {
			r.characterDirection = objects.CharacterDirection_RIGHT
		}
	} else {
		if r.y > targetY {
			r.characterDirection = objects.CharacterDirection_UP
		} else if r.y < targetY {
			r.characterDirection = objects.CharacterDirection_DOWN
		}
	}

	r.WalkAnimation.Update(r.characterDirection)
}
