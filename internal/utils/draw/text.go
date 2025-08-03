package draw

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/rxznik/GolangSnakeGame/internal/common"
)

type CalculateTextPositionFunc func(width, height float64) (x, y float64)

func CenterTextPosition(w, h float64) (x, y float64) {
	return float64(common.ScreenWidth-w) / 2, float64(common.ScreenHeight-h) / 2
}

var mplusFaceSource *text.GoTextFaceSource

func init() {
	var err error
	mplusFaceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
}

type TextOptions struct {
	FontSize              float64
	TextColor             color.Color
	CalculateTextPosition CalculateTextPositionFunc
}

func (opts *TextOptions) Parse() {
	if opts.TextColor == nil {
		opts.TextColor = common.TextColor
	}
	if opts.FontSize == 0 {
		opts.FontSize = common.FontSize
	}
	if opts.CalculateTextPosition == nil {
		opts.CalculateTextPosition = CenterTextPosition
	}
}

func DrawText(screen *ebiten.Image, message string, opts *TextOptions) {
	if opts == nil {
		opts = &TextOptions{}
	}
	opts.Parse()

	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   opts.FontSize,
	}

	w, h := text.Measure(message, face, face.Size)
	drawOpts := &text.DrawOptions{}
	drawOpts.GeoM.Translate(opts.CalculateTextPosition(w, h))
	drawOpts.ColorScale.ScaleWithColor(opts.TextColor)

	text.Draw(screen, message, face, drawOpts)
}
