package entity

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rxznik/GolangSnakeGame/internal/common"
	"github.com/rxznik/GolangSnakeGame/internal/utils/draw"
	"github.com/rxznik/GolangSnakeGame/internal/utils/point"
)

const PlayerTag = "player"

var playerColor = color.RGBA{R: 100, G: 255, B: 100, A: 255}

type Player struct {
	body      []point.Point
	direction point.Point
	latency   time.Duration
	score     int
}

func NewPlayer(initialBody []point.Point, direction point.Point) *Player {
	return &Player{
		body:      initialBody,
		direction: direction,
		latency:   common.GameSpeed,
	}
}

func (p *Player) Update(world worldView) bool {
	head := p.body[0]
	if head.IsInBounds(common.MaxX, common.MaxY) || head.IsInSlice(p.body[1:]) {
		return false
	}
	newHead := head.Copy()
	newHead.Add(p.direction)
	var grow bool
	for _, entity := range world.GetEntities(FoodTag) {
		food := entity.(*Food)
		if newHead.Equals(food.position) {
			grow = true
			food.Respawn()
		}
	}

	if grow {
		p.body = append([]point.Point{newHead}, p.body...)
		p.score++
		p.latency = time.Duration(float64(p.latency) * 0.99)
	} else {
		p.body = append([]point.Point{newHead}, p.body[:len(p.body)-1]...)
	}

	return true
}

func (p *Player) SetDirection(direction point.Point) {
	p.direction = direction
}

func (p Player) Latency() time.Duration {
	return p.latency
}

func (p *Player) Draw(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", p.score)

	draw.DrawText(screen, scoreText, &draw.TextOptions{
		CalculateTextPosition: func(w, h float64) (x, y float64) {
			return w, h
		},
	})

	for _, pt := range p.body {
		vector.DrawFilledRect(
			screen,
			float32(pt.X*common.GridSize),
			float32(pt.Y*common.GridSize),
			common.GridSize,
			common.GridSize,
			playerColor,
			true,
		)
	}
}

func (p *Player) Tag() string {
	return PlayerTag
}
