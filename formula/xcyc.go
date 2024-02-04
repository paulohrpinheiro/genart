package formula

import (
	"image"
	"math"
)

func XcycDraw(fs *FormulaStruct) {
	rect := image.Rect(0, 0, fs.ImageSize, fs.ImageSize)
	img := image.NewPaletted(rect, fs.Palette)

	constant := float64(fs.Constant)
	for x := range fs.ImageSize {
		xFloat := float64(x)
		for y := range fs.ImageSize {
			value := int(math.Pow(xFloat, constant) + math.Pow(float64(y), constant))
			color := uint8(value % fs.MaxColors)

			img.SetColorIndex(x, y, color)
		}
	}

	fs.Image = img
}
