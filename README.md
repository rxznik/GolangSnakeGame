# Simple Snake Game
Snake Game implementation in Golang using Ebiten Game Engine

## Features

* Eat food
* Change direction
* Game over
* Score
* Change speed based on score

## How to play

Download the game for your platform:

* [Windows](https://github.com/rxznik/GolangSnakeGame/bin/windows/)
* [Linux](https://github.com/rxznik/GolangSnakeGame/bin/linux/) (Comming soon)
* [macOS](https://github.com/rxznik/GolangSnakeGame/bin/macOS/)

Use arrow keys or WASD to move snake. Eat food to grow. Game over when snake hits itself or wall.



## How to run

Run the game using **go** command:
```bash
go run cmd/main.go
```

Or you can use **make**:
```bash
make run
```

## How to build

Build the game using **go** command:
```bash
go build cmd/main.go
./main
# With name
go build cmd/main.go -o SnakeGame
./SnakeGame
```

Or you can use **make**:
```bash
make build
./SnakeGame
# With name
make build DST=SnakeGame
./SnakeGame
```


