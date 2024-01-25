package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw05_shapes/shape"
)

func calculateArea(s any) (float64, error) {
	shape, ok := s.(shape.Shape)
	if !ok {
		return 0, fmt.Errorf("this is not a shape")
	}
	return shape.ShapeArea(), nil
}

func main() {
	var area float64
	var err error

	c1 := shape.Circle{
		Radius: 5,
	}

	fmt.Printf("Круг: радиус %d \r\n", c1.Radius)

	area, err = calculateArea(&c1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	r1 := shape.Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Printf("Прямоугольник: ширина %d, высота %d \r\n", r1.Width, r1.Height)

	area, err = calculateArea(&r1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	t1 := shape.Triangle{
		Base:   8,
		Height: 6,
	}

	fmt.Printf("Треугольник: основание %d, высота %d \r\n", t1.Base, t1.Height)

	area, err = calculateArea(&t1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	type Dog struct {
		weight int
	}

	d1 := Dog{
		weight: 12,
	}

	area, err = calculateArea(&d1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}
}
