package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/objects"
	"github.com/willcliffy/keydream/client/world/models"
)

type Character interface {
	Type() objects.CharacterType

	Update()
	Draw(screen *ebiten.Image)

	HandleMessage(msg world_models.WorldMessage)
}

