package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rxznik/GolangSnakeGame/internal/common"
	"github.com/rxznik/GolangSnakeGame/internal/utils/point"
)

const FoodTag = "food"

var foodColor = color.RGBA{R: 255, G: 100, B: 100, A: 255}

type Food struct {
	position point.Point
}

func NewFood() *Food {
	return &Food{
		position: point.RandomPoint(common.MaxX, common.MaxY),
	}
}

func (f *Food) Respawn() {
	f.position = point.RandomPoint(common.MaxX, common.MaxY)
}

func (f *Food) Update(world worldView) bool {
	return true
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(f.position.X*common.GridSize),
		float32(f.position.Y*common.GridSize),
		common.GridSize,
		common.GridSize,
		foodColor,
		true,
	)
}

func (f *Food) Tag() string {
	return FoodTag
}
