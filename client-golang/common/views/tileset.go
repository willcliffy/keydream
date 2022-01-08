package views

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/common/constants"
)

const (
	TilesetWidthInTiles = 16
	TilesetHeightInTiles = 16

	GrassBoundaryX0 = 0
	GrassBoundaryX1 = 8
	GrassBoundaryY0 = 0
	GrassBoundaryY1 = 8

	FlowerBoundaryX0 = 8
	FlowerBoundaryX1 = TilesetWidthInTiles
	FlowerBoundaryY0 = 0
	FlowerBoundaryY1 = 8

	TileBoundaryX0 = 0
	TileBoundaryX1 = 7
	TileBoundaryY0 = 8
	TileBoundaryY1 = TilesetHeightInTiles - 1
)

type Tileset struct {
	Grass []*ebiten.Image
	Flowers []*ebiten.Image
	Tiles []*ebiten.Image
}

func NewItemIcon(item string) (*ebiten.Image, error) {
	f, err := os.Open("./assets/objects/oliver_crawford/"+item+"_icon.png")
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func NewTileset() (*Tileset, error) {
	f, err := os.Open("./assets/environment/cainos/TX Tileset Grass.png")
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	var tileset Tileset
	wholeTileset := ebiten.NewImageFromImage(img)

	for x := 0; x < TilesetWidthInTiles; x++ {
		for y := 0; y < TilesetHeightInTiles; y++ {
			if x >= GrassBoundaryX0 && x < GrassBoundaryX1 && y >= GrassBoundaryY0 && y < GrassBoundaryY1 {
				tileset.Grass = append(tileset.Grass, wholeTileset.SubImage(
					image.Rect(
						x * constants.TileSize,
						y * constants.TileSize,
						(x + 1) * constants.TileSize,
						(y + 1) * constants.TileSize)).(*ebiten.Image))
			} else if x >= FlowerBoundaryX0 && x < FlowerBoundaryX1 && y >= FlowerBoundaryY0 && y < FlowerBoundaryY1 {
				tileset.Flowers = append(tileset.Flowers, wholeTileset.SubImage(
					image.Rect(
						x * constants.TileSize,
						y * constants.TileSize,
						(x + 1) * constants.TileSize,
						(y + 1) * constants.TileSize)).(*ebiten.Image))
			} else if x >= TileBoundaryX0 && x < TileBoundaryX1 && y >= TileBoundaryY0 && y < TileBoundaryY1 {
				// quirk of this particular tileset - random patch of grass :(
				if x == 4 {
					continue
				}

				tileset.Tiles = append(tileset.Tiles, wholeTileset.SubImage(
					image.Rect(
						x * constants.TileSize,
						y * constants.TileSize,
						(x + 1) * constants.TileSize,
						(y + 1) * constants.TileSize)).(*ebiten.Image))
			}
		}
	}

	return &tileset, nil
}

func (this *Tileset) Draw(screen *ebiten.Image) {
	for x := 0; x < constants.ScreenWidthInTiles; x++ {
		for y := 0; y < constants.ScreenHeightInTiles; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x * constants.TileSize), float64(y * constants.TileSize))

			// This is only necessary because I have an enormous screen.
			// todo - One day I'll make a settings menu and make this a setting.
			op.GeoM.Scale(2, 2)

			screen.DrawImage(this.Tiles[0], op)
		}
	}
}

// used for debug purposes to show groups of tiles.
func (this *Tileset) Display(screen *ebiten.Image) {
	x := 1
	y := 1
	for _, grassTile := range this.Grass {
		if y >= TilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * constants.TileSize), float64(y * constants.TileSize))
		screen.DrawImage(grassTile, op)
		y++
	}

	x = 10
	y = 1

	for _, flowerTile := range this.Flowers {
		if y >= TilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * constants.TileSize), float64(y * constants.TileSize))
		screen.DrawImage(flowerTile, op)
		y++
	}

	x = 20
	y = 1

	for _, tile := range this.Tiles {
		if y >= TilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * constants.TileSize), float64(y * constants.TileSize))
		screen.DrawImage(tile, op)
		y++
	}
}
