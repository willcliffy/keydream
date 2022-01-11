package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/objects"
)

type Character interface {
	Update()
	Draw(screen *ebiten.Image)
	HandleMessage(msg objects.WorldMessage)
}

