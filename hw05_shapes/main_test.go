package main

import (
	"testing"

	"github.com/TOsmanov/my-hw/hw05_shapes/shape"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var area float64
	var err error
	c1 := shape.Circle{
		Radius: 3,
	}
	area, err = calculateArea(&c1)
	assert.Equal(t, 28.274333882308138, area)
	assert.Nil(t, err)

	r1 := shape.Rectangle{
		Width:  8,
		Height: 7,
	}
	area, err = calculateArea(&r1)
	assert.Equal(t, float64(56), area)
	assert.Nil(t, err)

	t1 := shape.Triangle{
		Base:   3,
		Height: 15,
	}
	area, err = calculateArea(&t1)
	assert.Equal(t, float64(22), area)
	assert.Nil(t, err)

	cat1 := struct {
		weight int
	}{
		weight: 15,
	}
	area, err = calculateArea(&cat1)
	assert.Equal(t, float64(0), area)
	assert.NotNil(t, err)
}
