package printer

import (
	"testing"

	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestPrintStaff(t *testing.T) {
	data := []types.Employee{
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
	expected := "User ID: 2, Age: 45, Name: Tom, Department ID: 51; User ID: 23, Age: 31, Name: Jerry, Department ID: 346; "
	assert.Equal(t, expected, PrintStaff(data))
}
