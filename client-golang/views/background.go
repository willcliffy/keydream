package views

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/willcliffy/keydream/client/utils"
)

const (
	BackgroundTilesetWidthInTiles = 16
	BackgroundTilesetHeightInTiles = 16

	GrassBoundaryX0 = 0
	GrassBoundaryX1 = 8
	GrassBoundaryY0 = 0
	GrassBoundaryY1 = 8

	FlowerBoundaryX0 = 8
	FlowerBoundaryX1 = BackgroundTilesetWidthInTiles
	FlowerBoundaryY0 = 0
	FlowerBoundaryY1 = 8

	TileBoundaryX0 = 0
	TileBoundaryX1 = 7
	TileBoundaryY0 = 8
	TileBoundaryY1 = BackgroundTilesetHeightInTiles - 1
)

type Background struct {
	Grass []*ebiten.Image
	Flowers []*ebiten.Image
	Tiles []*ebiten.Image
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

	var background Background
	tileset := ebiten.NewImageFromImage(img)

	for x := 0; x < BackgroundTilesetWidthInTiles; x++ {
		for y := 0; y < BackgroundTilesetHeightInTiles; y++ {
			if x >= GrassBoundaryX0 && x < GrassBoundaryX1 && y >= GrassBoundaryY0 && y < GrassBoundaryY1 {
				background.Grass = append(background.Grass, tileset.SubImage(
					image.Rect(
						x * utils.TileSize,
						y * utils.TileSize,
						(x + 1) * utils.TileSize,
						(y + 1) * utils.TileSize)).(*ebiten.Image))
			} else if x >= FlowerBoundaryX0 && x < FlowerBoundaryX1 && y >= FlowerBoundaryY0 && y < FlowerBoundaryY1 {
				background.Flowers = append(background.Flowers, tileset.SubImage(
					image.Rect(
						x * utils.TileSize,
						y * utils.TileSize,
						(x + 1) * utils.TileSize,
						(y + 1) * utils.TileSize)).(*ebiten.Image))
			} else if x >= TileBoundaryX0 && x < TileBoundaryX1 && y >= TileBoundaryY0 && y < TileBoundaryY1 {
				// quirk of this particular tileset - random patch of grass :(
				if x == 4 {
					continue
				}

				background.Tiles = append(background.Tiles, tileset.SubImage(
					image.Rect(
						x * utils.TileSize,
						y * utils.TileSize,
						(x + 1) * utils.TileSize,
						(y + 1) * utils.TileSize)).(*ebiten.Image))
			}
		}
	}

	return &background, nil
}

func (this *Background) Draw(screen *ebiten.Image) {
	for x := 0; x < utils.ScreenWidthInTiles; x++ {
		for y := 0; y < utils.ScreenHeightInTiles; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x * utils.TileSize), float64(y * utils.TileSize))
			// This is only necessary because I have an enormous screen.
			// todo - One day I'll make a settings menu and make this a setting.
			op.GeoM.Scale(2, 2)
			screen.DrawImage(this.Tiles[0], op)
		}
	}
}

// used for debug purposes to show groups of tiles.
func (this *Background) Display(screen *ebiten.Image) {
	x := 1
	y := 1
	for _, grassTile := range this.Grass {
		if y >= BackgroundTilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * utils.TileSize), float64(y * utils.TileSize))
		screen.DrawImage(grassTile, op)
		y++
	}

	x = 10
	y = 1

	for _, flowerTile := range this.Flowers {
		if y >= BackgroundTilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * utils.TileSize), float64(y * utils.TileSize))
		screen.DrawImage(flowerTile, op)
		y++
	}

	x = 20
	y = 1

	for _, tile := range this.Tiles {
		if y >= BackgroundTilesetHeightInTiles-1 {
			x++
			y = 1
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x * utils.TileSize), float64(y * utils.TileSize))
		screen.DrawImage(tile, op)
		y++
	}
}
