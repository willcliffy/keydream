package objects

import "github.com/hajimehoshi/ebiten/v2"

type CharacterType int
const (
	CharacterType_LOCAL  CharacterType = 0
	CharacterType_REMOTE CharacterType = 1
	CharacterType_NPC    CharacterType = 2
)

func CharacterType_values() []CharacterType {
	return []CharacterType{
		CharacterType_LOCAL,
		CharacterType_REMOTE,
		CharacterType_NPC,
	}
}

func (this CharacterType) String() string {
	switch this {
	case CharacterType_LOCAL:
		return "local"
	case CharacterType_REMOTE:
		return "remote"
	case CharacterType_NPC:
		return "npc"
	}
	return ""
}


type CharacterState int
const (
	CharacterState_IDLE CharacterState = 0
	CharacterState_WALK CharacterState = 1
)

func CharacterState_values() []CharacterState {
	return []CharacterState{
		CharacterState_IDLE,
		CharacterState_WALK,
	}
}

func (this CharacterState) String() string {
	switch this {
	case CharacterState_IDLE:
		return "idle"
	case CharacterState_WALK:
		return "walk"
	}
	return ""
}


type CharacterDirection ebiten.Key
const (
	CharacterDirection_NONE  CharacterDirection = 0
	CharacterDirection_UP    CharacterDirection = 1
	CharacterDirection_DOWN  CharacterDirection = 2
	CharacterDirection_LEFT  CharacterDirection = 3
	CharacterDirection_RIGHT CharacterDirection = 4
)

func (this CharacterDirection) ContainedIn(directions []CharacterDirection) bool {
	for _, dir := range directions {
		if dir == this {
			return true
		}
	}
	return false
}

func (this CharacterDirection) OppositeContainedIn(direction []CharacterDirection) bool {
	for _, dir := range direction {
		if dir.Opposite() == this {
			return true
		}
	}

	return false
}

func (this CharacterDirection) Opposite() CharacterDirection {
	switch this {
	case CharacterDirection_UP:
		return CharacterDirection_DOWN
	case CharacterDirection_DOWN:
		return CharacterDirection_UP
	case CharacterDirection_LEFT:
		return CharacterDirection_RIGHT
	case CharacterDirection_RIGHT:
		return CharacterDirection_LEFT
	}
	return CharacterDirection_NONE
}

func CharacterDirection_values() []CharacterDirection {
	return []CharacterDirection{
		CharacterDirection_UP,
		CharacterDirection_DOWN,
		CharacterDirection_LEFT,
		CharacterDirection_RIGHT,
	}
}

func KeyToDirection(key ebiten.Key) CharacterDirection {
	switch key {
	case ebiten.KeyUp, ebiten.KeyW:
		return CharacterDirection_UP
	case ebiten.KeyDown, ebiten.KeyS:
		return CharacterDirection_DOWN
	case ebiten.KeyLeft, ebiten.KeyA:
		return CharacterDirection_LEFT
	case ebiten.KeyRight, ebiten.KeyD:
		return CharacterDirection_RIGHT
	}
	return CharacterDirection_NONE
}


func (this CharacterDirection) String() string {
	switch this {
	case CharacterDirection_UP:
		return "up"
	case CharacterDirection_DOWN:
		return "down"
	case CharacterDirection_LEFT:
		return "left"
	case CharacterDirection_RIGHT:
		return "right"
	}
	return ""
}
