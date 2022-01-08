package objects

import "github.com/hajimehoshi/ebiten/v2"

type CharacterType int
const (
	CharacterType_Local CharacterType = iota
	CharacterType_Remote
	CharacterType_NPC
)

func CharacterType_values() []CharacterType {
	return []CharacterType{
		CharacterType_Local,
		CharacterType_Remote,
		CharacterType_NPC,
	}
}

func (this CharacterType) String() string {
	switch this {
	case CharacterType_Local:
		return "local"
	case CharacterType_Remote:
		return "remote"
	case CharacterType_NPC:
		return "npc"
	}
	return ""
}


type CharacterState int
const (
	CharacterState_Idle CharacterState = iota
	CharacterState_Walk
)

func CharacterState_values() []CharacterState {
	return []CharacterState{
		CharacterState_Idle,
		CharacterState_Walk,
	}
}

func (this CharacterState) String() string {
	switch this {
	case CharacterState_Idle:
		return "idle"
	case CharacterState_Walk:
		return "walk"
	}
	return ""
}


type CharacterDirection ebiten.Key
const (
	CharacterDirection_Up    = CharacterDirection(ebiten.KeyUp)
	CharacterDirection_Down  = CharacterDirection(ebiten.KeyDown)
	CharacterDirection_Left  = CharacterDirection(ebiten.KeyLeft)
	CharacterDirection_Right = CharacterDirection(ebiten.KeyRight)
)

func CharacterDirection_values() []CharacterDirection {
	return []CharacterDirection{
		CharacterDirection_Up,
		CharacterDirection_Down,
		CharacterDirection_Left,
		CharacterDirection_Right,
	}
}

func (this CharacterDirection) String() string {
	switch this {
	case CharacterDirection_Up:
		return "up"
	case CharacterDirection_Down:
		return "down"
	case CharacterDirection_Left:
		return "left"
	case CharacterDirection_Right:
		return "right"
	}
	return ""
}
