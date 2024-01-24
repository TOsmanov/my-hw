package reader

import (
	"testing"

	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	var data []types.Employee
	var err error

	// Case 1
	data, err = ReadJSON("not-exist.json")

	assert.Nil(t, data)
	assert.NotNil(t, err)

	// Case 2
	data, err = ReadJSON("/proc/self/mem")
	t.Log(data, err)

	assert.Nil(t, data)
	assert.NotNil(t, err)

	// Case 3
	data, err = ReadJSON("test-data.json")
	expected := []types.Employee{
		{
			UserID:       2,
			Age:          45,
			Name:         "Tom",
			DepartmentID: 51,
		},
		{
			UserID:       23,
			Age:          31,
			Name:         "Jerry",
			DepartmentID: 346,
		},
	}

	assert.Equal(t, data, expected)
	assert.Nil(t, err)
}
