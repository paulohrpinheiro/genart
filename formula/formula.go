package formula

import (
	"errors"
	"image"
	"image/color"
)

type PaletteType = color.Palette

type DrawFuncType func(*FormulaStruct)

type FormulaStruct struct {
	ImageSize int
	Constant  int
	MaxColors int
	Palette   PaletteType
	Image	  image.Image
	DrawFunc  func(*FormulaStruct)
}

func New() *FormulaStruct {
	return &FormulaStruct{}
}

type fn func(*FormulaStruct)

func functionByName(name string) (fn, error) {
	formulasMap := map[string]fn{
		"ccxy": CcxyDraw,
		"xcyc": XcycDraw,
	}

	newFunction := formulasMap[name]
	if newFunction == nil {
		return nil, errors.New("Unknown formula.")
	}

	return newFunction, nil
}

func (fs *FormulaStruct) Init(
	functionName string,
	imageSize int,
	constant int,
	maxColors int,
	palette PaletteType,
) error {
	var err error

	fs.ImageSize = imageSize
	fs.Constant = constant
	fs.MaxColors = maxColors
	fs.Palette = palette

	fs.DrawFunc, err = functionByName(functionName)
	return err
}

func (fs *FormulaStruct) Draw() {
	fs.DrawFunc(fs)
}
