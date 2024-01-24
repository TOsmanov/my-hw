package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	e := Employee{
		UserID:       1,
		Age:          2,
		Name:         "Rob",
		DepartmentID: 3,
	}

	assert.Equal(t, e.String(), "User ID: 1, Age: 2, Name: Rob, Department ID: 3; ")
}
