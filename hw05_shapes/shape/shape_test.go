package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	c1 := Circle{
		Radius: 9,
	}
	assert.Equal(t, 254.46900494077323, c1.ShapeArea())
}

func TestRectangle(t *testing.T) {
	r1 := Rectangle{
		Width:  12,
		Height: 8,
	}
	assert.Equal(t, float64(96), r1.ShapeArea())
}

func TestTriangle(t *testing.T) {
	t1 := Triangle{
		Base:   5,
		Height: 12,
	}
	assert.Equal(t, float64(30), t1.ShapeArea())
}
