package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		desc     string
		list     []int
		key      int
		expected int
	}{
		{
			desc:     "Case 1",
			list:     []int{-5, 1, 4, 5, 87, 100},
			key:      87,
			expected: 4,
		},
		{
			desc:     "Case 2",
			list:     []int{45, 46, 54, 123, 257, 1568, 1600, 1857, 2000},
			key:      54,
			expected: 2,
		},
		{
			desc:     "Case 3",
			list:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			key:      5,
			expected: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, found := BinarySearch(tC.list, tC.key)
			assert.True(t, found)
			assert.Equal(t, tC.expected, result)
		})
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	list := []int{1, 2, 6, 50, 54, 123, 257, 1568, 1600, 1857, 2000}
	key := 256
	result, found := BinarySearch(list, key)

	assert.False(t, found)
	assert.Equal(t, -1, result)
}
