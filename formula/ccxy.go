package formula

import (
	"image"
)

func CcxyDraw(fs *FormulaStruct) {
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
