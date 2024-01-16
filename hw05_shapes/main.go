package main

import (
	"fmt"

	. "github.com/TOsmanov/my-hw/hw05_shapes/shape"
)

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("this is not a shape")
	}
	return shape.ShapeArea(), nil
}

func main() {
	var area float64
	var err error

	var c1 = Circle{
		Radius: 5,
	}

	area, err = calculateArea(&c1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	var r1 = Rectangle{
		Width:  10,
		Height: 5,
	}

	area, err = calculateArea(&r1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	var t1 = Triangle{
		Base:   8,
		Height: 6,
	}

	area, err = calculateArea(&t1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}

	type Dog struct {
		weight int
	}

	var d1 = Dog{
		weight: 12,
	}

	area, err = calculateArea(&d1)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
	} else {
		fmt.Printf("Площадь: %v\r\n\n", area)
	}
}
