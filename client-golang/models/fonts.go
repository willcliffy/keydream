package models

import (
	"github.com/willcliffy/keydream/client/assets/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontSize int

const (
	FontSizeTiny FontSize   = 8
	FontSizeSmall FontSize  = 12
	FontSizeMedium FontSize = 16
	FontSizeLarge FontSize  = 32
	FontSizeHuge FontSize   = 64
	FontSizeMax FontSize    = FontSizeHuge
)

func LoadFonts() (map[FontSize]font.Face, error) {
	sfntFont, err := opentype.Parse(fonts.LunchDS_ttf)
	if err != nil {
		return nil, nil
	}

	fonts := make(map[FontSize]font.Face)

	for _, size := range []FontSize{
		FontSizeTiny,
		FontSizeSmall,
		FontSizeMedium,
		FontSizeLarge,
		FontSizeHuge,
	} {
		fonts[size], err = opentype.NewFace(sfntFont, &opentype.FaceOptions{
			Size: float64(size),
			DPI: 144,
			Hinting: font.HintingFull,
		})
		if err != nil {
			return nil, err
		}
	}

	return fonts, nil
}
