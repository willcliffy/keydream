package models

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	FrameDuration time.Duration

	frames []*ebiten.Image
	currentFrame  int
	lastFrameTime time.Time
}

func NewAnimation(
	frames []image.Image,
	frameDuration time.Duration,
) *Animation {
	framesCopy := make([]*ebiten.Image, len(frames))
	for i, frame := range frames {
		framesCopy[i] = ebiten.NewImageFromImage(frame)
	}

	return &Animation{
		FrameDuration: frameDuration,

		frames: framesCopy,
		currentFrame: 0,
		lastFrameTime: time.Now(),
	}
}

func (this *Animation) Update() {
	if time.Since(this.lastFrameTime) > this.FrameDuration {
		this.currentFrame = (this.currentFrame + 1) % len(this.frames)
		this.lastFrameTime = time.Now()
	}
}

func (this Animation) GetCurrentFrame() *ebiten.Image {
	return this.frames[this.currentFrame]
}

func (this Animation) Draw(screen *ebiten.Image) {
	screen.DrawImage(this.GetCurrentFrame(), nil)
}
