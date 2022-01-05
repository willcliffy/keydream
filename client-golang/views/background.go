package views

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/utils"
)

type Background struct {
	Tiles map[int]map[int]*ebiten.Image
}

func NewBackground() (*Background, error) {
	f, err := os.Open("./assets/environment/cainos/TX Tileset Grass.png")
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	tileset := ebiten.NewImageFromImage(img)

	var background Background
	background.Tiles = make(map[int]map[int]*ebiten.Image)

	for x := 0; x < utils.ScreenWidthInTiles; x++ {
		background.Tiles[x] = make(map[int]*ebiten.Image)
		for y := 0; y < utils.ScreenHeightInTiles; y++ {
			background.Tiles[x][y] = tileset.SubImage(
				image.Rect(
					x * utils.TileWidth,
					y * utils.TileHeight,
					x * utils.TileWidth + utils.TileWidth,
					y * utils.TileHeight + utils.TileHeight)).(*ebiten.Image)
		}
	}

	return &background, nil
}

func (this *Background) Draw(screen *ebiten.Image) {
	for x := 0; x < utils.ScreenWidthInTiles; x++ {
		for y := 0; y < utils.ScreenHeightInTiles; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x * utils.TileWidth), float64(y * utils.TileHeight))
			screen.DrawImage(this.Tiles[x][y], op)
		}
	}
}
