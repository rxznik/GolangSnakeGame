package game

import (
	"errors"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rxznik/GolangSnakeGame/internal/common"
	"github.com/rxznik/GolangSnakeGame/internal/entity"
	"github.com/rxznik/GolangSnakeGame/internal/utils/draw"
	"github.com/rxznik/GolangSnakeGame/internal/utils/point"
	"github.com/rxznik/GolangSnakeGame/internal/world"
)

var (
	directionUp    = point.Point{X: 0, Y: -1}
	directionDown  = point.Point{X: 0, Y: 1}
	directionLeft  = point.Point{X: -1, Y: 0}
	directionRight = point.Point{X: 1, Y: 0}
)

type Game struct {
	world      *world.World
	lastUpdate time.Time
	gameOver   bool
}

func New() *Game {
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Golang Snake Game")
	g := new(Game)
	g.Restart()
	return g
}

func (g *Game) SetupWorld() {
	g.world = world.New()
	g.world.AddEntity(entity.NewPlayer(
		[]point.Point{
			{X: common.MaxX / 2, Y: common.MaxY / 2},
			{X: common.MaxX/2 - 1, Y: common.MaxY / 2},
		},
		directionRight,
	))
	g.world.AddEntity(entity.NewFood())
}

func (g *Game) Restart() {
	g.SetupWorld()
	g.gameOver = false
	g.lastUpdate = time.Now()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}

func (g *Game) Update() error {

	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Restart()
		}
		return nil
	}

	playerRaw := g.world.GetFirstEntity(entity.PlayerTag)
	if playerRaw == nil {
		return errors.New("player not found")
	}
	player := playerRaw.(*entity.Player)

	g.handlePlayerKeyPress(player)

	if time.Since(g.lastUpdate) < player.Latency() {
		return nil
	}

	for _, e := range g.world.Entities() {
		if !e.Update(g.world) {
			g.gameOver = true
			return nil
		}
	}

	g.lastUpdate = time.Now()
	return nil
}

func (g *Game) handlePlayerKeyPress(player *entity.Player) {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		player.SetDirection(directionUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		player.SetDirection(directionDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.SetDirection(directionLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.SetDirection(directionRight)
	}
}

const gameOverFontSize = 32

var gameOverColor = color.RGBA{R: 255, G: 50, B: 250, A: 255}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		draw.DrawText(screen, "Game Over", &draw.TextOptions{
			FontSize:  gameOverFontSize,
			TextColor: gameOverColor,
		})
		draw.DrawText(screen, "Press Space to restart", &draw.TextOptions{
			CalculateTextPosition: func(w, h float64) (x, y float64) {
				x, y = draw.CenterTextPosition(w, h)
				return x, y + common.GridSize*2
			},
		})
	}

	for _, e := range g.world.Entities() {
		e.Draw(screen)
	}
}
