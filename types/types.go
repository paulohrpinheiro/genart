package genart

import (
	"image"
	"image/color"
)

type PaletteType = color.Palette

type FormulaStruct struct {
	ImageSize int
	Constant  int
	MaxColors int
	Palette   PaletteType
	Image	  image.Image
}

type FormulaInterface interface {
	Init(imageSize, constant int, maxColors int, palette PaletteType)
	Draw()
}
