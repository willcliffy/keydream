package world

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
)

type RemoteCharacter struct {
	Character

	Name string
	ID uint8

	animation *views.CharacterAnimation

	x, y float64
	speedX, speedY float64

	targets []objects.Position
}

func NewRemoteCharacter(id uint8, characterName string, x, y float64) *RemoteCharacter {
	animation := views.NewCharacterAnimation()
	return &RemoteCharacter{
		Name:      characterName,
		ID:        id,

		animation: &animation,

		x:         x,
		y:         y,
		speedX:    0,
		speedY:    0,
		targets:   []objects.Position{},
	}
}

func (r *RemoteCharacter) Update() {
	if len(r.targets) == 0 {
		r.animation.Update()
		return
	}

	directions := r.directionsToTarget(r.targets[0])
	for _, direction := range directions {
		r.moveInDirection(direction)
		r.animation.Walk(direction)
	}

	reachedX := r.reachedTargetX()
	reachedY := r.reachedTargetY()

	if reachedX {
		r.x = r.targets[0].X
		r.speedX = 0
	}

	if reachedY {
		r.y = r.targets[0].Y
		r.speedY = 0
	}

	if reachedX && reachedY {
		r.targets = r.targets[1:]
		r.animation.Idle()
	}

	r.animation.Update()
}

func (r *RemoteCharacter) Draw(screen *ebiten.Image) {
	r.animation.Draw(screen, r.x, r.y)
}

func (r *RemoteCharacter) HandleMessage(msg *objects.WorldMessage) error {
	switch msg.Type {
	case objects.WorldMessageType_MOVE:
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

		fmt.Printf("%s moved to %f, %f\n", r.Name, x, y)
		r.targets = append(r.targets, objects.Position{X: x, Y: y})
		return nil
	default:
		return fmt.Errorf("unknown message type for remote character: %s", msg.Type)
	}
}

func (r *RemoteCharacter) moveInDirection(direction objects.CharacterDirection) {
	switch direction {
	case objects.CharacterDirection_LEFT:
		r.x -= constants.RemoteCharacterMaxWalkSpeed
	case objects.CharacterDirection_RIGHT:
		r.x += constants.RemoteCharacterMaxWalkSpeed
	case objects.CharacterDirection_UP:
		r.y -= constants.RemoteCharacterMaxWalkSpeed
	case objects.CharacterDirection_DOWN:
		r.y += constants.RemoteCharacterMaxWalkSpeed
	}
}

func (r RemoteCharacter) directionsToTarget(target objects.Position) []objects.CharacterDirection {
	ret := []objects.CharacterDirection{}
	if target.X < r.x + constants.RemoteCharacterAlpha {
		ret = append(ret, objects.CharacterDirection_LEFT)
	} else if target.X > r.x + constants.RemoteCharacterAlpha {
		ret = append(ret, objects.CharacterDirection_RIGHT)
	}

	if target.Y < r.y - constants.RemoteCharacterAlpha {
		ret = append(ret, objects.CharacterDirection_UP)
	} else if target.Y > r.y + constants.RemoteCharacterAlpha {
		ret = append(ret, objects.CharacterDirection_DOWN)
	}

	return ret
}

func (r RemoteCharacter) reachedTargetX() bool {
	if len(r.targets) == 0 {
		return true
	}

	return math.Abs(r.x - r.targets[0].X) <= constants.RemoteCharacterAlpha
}

func (r RemoteCharacter) reachedTargetY() bool {
	if len(r.targets) == 0 {
		return true
	}

	return math.Abs(r.y - r.targets[0].Y) <= constants.RemoteCharacterAlpha	
}
