package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncremetation(t *testing.T) {
	counter := Counter{}
	counter.c = 10

	counter.incremetation(100)

	assert.Equal(t, counter.c, 110)
}
