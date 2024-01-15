package main

import (
	"flag"
	"image/png"
	"image/color/palette"
	"os"

	"genart/ccxy"
)

const (
	defaultImageSize = 1000
	defaultMaxColors = 150
	defaultConstant  = 10
	defaultFileName  = "formulaimg.png"
)

func main() {
	colors := palette.WebSafe

	fileName := flag.String("filename", defaultFileName, "filename to save image (.png)")
	imageSize := flag.Int("size", defaultImageSize, "size of image")
	constant := flag.Int("constant", defaultConstant, "constant for formula (c): c-c*x*y")
	maxColors := flag.Int("colors", len(colors), "number of colors")

	flag.Parse()

	newImage := ccxy.CcxyStruct{}
	newImage.Init(*imageSize, *constant, *maxColors, colors)
	newImage.Draw()

	f, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, newImage.Image)
}
