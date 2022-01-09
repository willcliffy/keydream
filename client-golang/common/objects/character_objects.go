package objects

import "github.com/hajimehoshi/ebiten/v2"

type CharacterType int
const (
	CharacterType_LOCAL CharacterType = iota
	CharacterType_REMOTE
	CharacterType_NPC
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
	CharacterState_IDLE CharacterState = iota
	CharacterState_WALK
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
	CharacterDirection_NONE CharacterDirection = iota
	CharacterDirection_UP
	CharacterDirection_DOWN
	CharacterDirection_LEFT
	CharacterDirection_RIGHT
)

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
