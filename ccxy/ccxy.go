package ccxy

import (
	"image"

	"genart/types"
)

type CcxyStruct genart.FormulaStruct

func (fs *CcxyStruct) Init(
	imageSize,
	constant int,
	maxColors int,
	palette genart.PaletteType,
) {
	fs.ImageSize = imageSize
	fs.Constant = constant
	fs.MaxColors = maxColors
	fs.Palette = palette
}

func (fs *CcxyStruct) Draw() {
	rect := image.Rect(0, 0, fs.ImageSize, fs.ImageSize)
	img := image.NewPaletted(rect, fs.Palette)

	for x := range fs.ImageSize {
		for y := range fs.ImageSize {
			value := (fs.Constant - fs.Constant*x*y)
			color := uint8(value % fs.MaxColors)

			img.SetColorIndex(x, y, color)
		}
	}

	fs.Image = img
}
