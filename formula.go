package main

import (
	"flag"
	"image"
	"image/color/palette"
	"image/png"
	"os"
)

const (
	defaultImageSize = 1000
	defaultMaxColors = 150
	defaultConstant  = 10
	defaultFileName  = "formula.png"
)

func main() {
	fileName := flag.String("filename", defaultFileName, "filename to save image (.png)")
	imageSize := flag.Int("size", defaultImageSize, "size of image")
	constant := flag.Int("constant", defaultConstant, "constant for formula (c): c-c*x*y")
	maxColors := flag.Int("colors", len(palette.WebSafe), "number of colors")

	flag.Parse()

	rect := image.Rect(0, 0, *imageSize, *imageSize)
	img := image.NewPaletted(rect, palette.WebSafe)

	for x := range *imageSize {
		for y := range *imageSize {
			value := (*constant - *constant*x*y)
			color := uint8(value % *maxColors)

			img.SetColorIndex(x, y, color)
		}
	}

	f, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, img)
}
