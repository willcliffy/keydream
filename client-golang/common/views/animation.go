package views

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames []*ebiten.Image
	CurrentFrame  int
	FrameDuration time.Duration
	LastFrameTime time.Time
}

func NewAnimation(frames []*ebiten.Image, frameDuration time.Duration) *Animation {
	return &Animation{
		Frames: frames,
		FrameDuration: frameDuration,
		LastFrameTime: time.Now(),
	}
}

func (this *Animation) Update() {
	if time.Since(this.LastFrameTime) > this.FrameDuration {
		this.CurrentFrame = (this.CurrentFrame + 1) % len(this.Frames)
		this.LastFrameTime = time.Now()
	}
}

func (this *Animation) GetCurrentFrame() *ebiten.Image {
	return this.Frames[this.CurrentFrame]
}
