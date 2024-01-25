package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestPrintStaff(t *testing.T) {
	var output bytes.Buffer

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
	a := "User ID: 2, Age: 45, Name: Tom, Department ID: 51; \n"
	b := "User ID: 23, Age: 31, Name: Jerry, Department ID: 346; \n"
	expected := fmt.Sprintf("%s%s", a, b)
	PrintStaff(&output, data)
	assert.Equal(t, expected, output.String())
}
