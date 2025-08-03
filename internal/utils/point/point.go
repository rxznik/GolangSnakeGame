package point

import (
	"math/rand"
	"slices"
)

type Point struct {
	X, Y int
}

func RandomPoint(maxX, maxY int) Point {
	return Point{
		X: rand.Intn(maxX),
		Y: rand.Intn(maxY),
	}
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) Add(other Point) {
	p.X += other.X
	p.Y += other.Y
}

func (p *Point) Subtract(other Point) {
	p.X -= other.X
	p.Y -= other.Y
}

func (p Point) Copy() Point {
	return Point{X: p.X, Y: p.Y}
}

func (p Point) IsInBounds(maxX, maxY int) bool {
	return p.X < 0 || p.X >= maxX || p.Y < 0 || p.Y >= maxY
}

func (p Point) IsInSlice(points []Point) bool {
	return slices.Contains(points, p)
}
