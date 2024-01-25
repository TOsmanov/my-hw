package shape

import (
	"math"
)

type Shape interface {
	ShapeArea() float64
}

type Circle struct {
	Radius int
}

func (c *Circle) ShapeArea() float64 {
	pi := math.Pi
	r := float64(c.Radius)
	return math.Pow(r, 2) * pi
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) ShapeArea() float64 {
	return float64(r.Width * r.Height)
}

type Triangle struct {
	Base   int
	Height int
}

func (t *Triangle) ShapeArea() float64 {
	return float64(t.Base * t.Height / 2)
}
