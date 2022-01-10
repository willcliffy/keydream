package views

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/objects"
)

type Animation struct {
	NumberOfFrames int
	FrameDuration time.Duration

	upFrames []*ebiten.Image
	downFrames []*ebiten.Image
	leftFrames []*ebiten.Image
	rightFrames []*ebiten.Image

	currentFrame  int
	lastFrameTime time.Time
}

func NewAnimation(
	up []*ebiten.Image,
	down []*ebiten.Image,
	left []*ebiten.Image,
	right []*ebiten.Image,
	frameDuration time.Duration,
) *Animation {
	if len(up) != len(down) || len(up) != len(left) || len(up) != len(right) {
		panic("All frames must be the same length")
	}

	return &Animation{
		NumberOfFrames: len(up),
		FrameDuration: frameDuration,

		upFrames: up,
		downFrames: down,
		leftFrames: left,
		rightFrames: right,

		currentFrame: 0,
		lastFrameTime: time.Now(),
	}
}

func (this *Animation) Update(direction objects.CharacterDirection) {
	if time.Since(this.lastFrameTime) > this.FrameDuration {
		this.currentFrame = (this.currentFrame + 1) % this.NumberOfFrames
		this.lastFrameTime = time.Now()
	}
}

func (this Animation) GetCurrentFrame(direction objects.CharacterDirection) *ebiten.Image {
	switch direction {
	case objects.CharacterDirection_UP:
		return this.upFrames[this.currentFrame]
	case objects.CharacterDirection_DOWN:
		return this.downFrames[this.currentFrame]
	case objects.CharacterDirection_LEFT:
		return this.leftFrames[this.currentFrame]
	case objects.CharacterDirection_RIGHT:
		return this.rightFrames[this.currentFrame]
	default:
		return nil
	}
}
