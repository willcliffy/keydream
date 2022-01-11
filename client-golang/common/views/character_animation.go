package views

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/willcliffy/keydream/client/common/constants"
	"github.com/willcliffy/keydream/client/common/objects"
)

type CharacterAnimation struct {
	State objects.CharacterState
	Direction objects.CharacterDirection

	currentFrame int
	frameTimer time.Time

	frames map[objects.CharacterDirection]map[objects.CharacterState][]*ebiten.Image
}

func NewCharacterAnimation() CharacterAnimation {
	return CharacterAnimation{
		State: objects.CharacterState_IDLE,
		Direction: objects.CharacterDirection_DOWN,

		currentFrame: 0,
		frameTimer: time.Now(),
		frames: initializeFrames(),
	}
}

func initializeFrames() map[objects.CharacterDirection]map[objects.CharacterState][]*ebiten.Image {
	ret := make(map[objects.CharacterDirection]map[objects.CharacterState][]*ebiten.Image)

	for _, direction := range objects.CharacterDirection_values() {
		ret[direction] = make(map[objects.CharacterState][]*ebiten.Image)
		for _, state := range objects.CharacterState_values() {
			ret[direction][state] = make([]*ebiten.Image, 4)
			for i := 0; i < 4; i++ {
				img, _, err := ebitenutil.NewImageFromFile(
					fmt.Sprintf("./assets/sprites/rgs_dev/Character without weapon/%s/%s %s%d.png", state.String(), state.String(), direction.String(), i+1))
				if err != nil {
					panic(err)
				}
				ret[direction][state][i] = img
			}
		}
	}

	return ret
}

func (this *CharacterAnimation) Update() {
	if time.Since(this.frameTimer) > constants.CharacterAnimationSpeed {
		this.currentFrame++
		if this.currentFrame >= len(this.frames[this.Direction][this.State]) {
			this.currentFrame = 0
		}
		this.frameTimer = time.Now()
	}
}

func (this *CharacterAnimation) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(constants.CharacterScale, constants.CharacterScale)

	screen.DrawImage(this.frames[this.Direction][this.State][this.currentFrame], op)
}

func (this *CharacterAnimation) Idle() {
	this.State = objects.CharacterState_IDLE
}

func (this *CharacterAnimation) Walk(direction objects.CharacterDirection) {
	this.State = objects.CharacterState_WALK
	this.Direction = direction
}
