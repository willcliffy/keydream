package world_utils

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/common/views"
)

func LoadIdleAnimations() *views.Animation {
	return views.NewAnimation(
		loadCharacterAnimation(objects.CharacterState_IDLE, objects.CharacterDirection_UP),
		loadCharacterAnimation(objects.CharacterState_IDLE, objects.CharacterDirection_DOWN),
		loadCharacterAnimation(objects.CharacterState_IDLE, objects.CharacterDirection_LEFT),
		loadCharacterAnimation(objects.CharacterState_IDLE, objects.CharacterDirection_RIGHT),
		time.Millisecond*250,
	)
}

func LoadWalkAnimations() *views.Animation {
	return views.NewAnimation(
		loadCharacterAnimation(objects.CharacterState_WALK, objects.CharacterDirection_UP),
		loadCharacterAnimation(objects.CharacterState_WALK, objects.CharacterDirection_DOWN),
		loadCharacterAnimation(objects.CharacterState_WALK, objects.CharacterDirection_LEFT),
		loadCharacterAnimation(objects.CharacterState_WALK, objects.CharacterDirection_RIGHT),
		time.Millisecond*250,
	)
}

func loadCharacterAnimation(state objects.CharacterState, direction objects.CharacterDirection) []*ebiten.Image {
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

	return frames
}
