package common

import (
	"image/color"
	"time"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
	GameSpeed    = 100 * time.Millisecond
)

const (
	MaxX = ScreenWidth / GridSize
	MaxY = ScreenHeight / GridSize
)

const FontSize float64 = 16

var TextColor = color.White
