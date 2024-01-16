package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeArea() float64
}

type Circle struct {
	Radius int
}

func (c *Circle) ShapeArea() float64 {
	fmt.Printf("Круг: радиус %d \r\n", c.Radius)
	pi := math.Pi
	r := float64(c.Radius)
	return math.Pow(r, 2) * pi
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) ShapeArea() float64 {
	fmt.Printf("Прямоугольник: ширина %d, высота %d \r\n", r.Width, r.Height)
	return float64(r.Width * r.Height)
}

type Triangle struct {
	Base   int
	Height int
}

func (t *Triangle) ShapeArea() float64 {
	fmt.Printf("Треугольник: основание %d, высота %d \r\n", t.Base, t.Height)
	return float64(t.Base * t.Height / 2)
}
