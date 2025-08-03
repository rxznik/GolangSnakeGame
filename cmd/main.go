package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rxznik/GolangSnakeGame/internal/game"
)

func main() {

	app := game.New()

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
