package main

import (
	"flag"
	"image/png"
	"image/color/palette"
	"os"

	"genart/formula"
)

const (
	defaultImageSize = 1000
	defaultMaxColors = 150
	defaultConstant  = 10
	defaultFileName  = "formulaimg.png"
)

func main() {
	colors := palette.WebSafe

	formulaName := flag.String("formula", "ccxy", "formula to use")
	fileName := flag.String("filename", defaultFileName, "filename to save image (.png)")
	imageSize := flag.Int("size", defaultImageSize, "size of image")
	constant := flag.Int("constant", defaultConstant, "constant for formula (c): c-c*x*y")
	maxColors := flag.Int("colors", len(colors), "number of colors")

	flag.Parse()

	newImage := formula.New()

	newImage.Init(*formulaName, *imageSize, *constant, *maxColors, colors)
	newImage.Draw()

	f, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, newImage.Image)
}
