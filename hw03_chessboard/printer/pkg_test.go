package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintLine(t *testing.T) {
	var expected string
	// Case 1
	expected = " # # "
	assert.Equal(t, expected, printLine(true, 5))
	// Case 2
	expected = "# # #"
	assert.Equal(t, expected, printLine(false, 5))
}

func TestPrintRows(t *testing.T) {
	expected := `+-----+
| # # |
|# # #|
| # # |
|# # #|
| # # |
+-----+
`
	assert.Equal(t, expected, PrintRows(5, 5))
}

func TestPrinter(t *testing.T) {
	testCases := []struct {
		desc            string
		input, expected string
	}{
		{
			desc:  "Case 1",
			input: "",
			expected: `+--------+
| # # # #|
|# # # # |
| # # # #|
|# # # # |
| # # # #|
|# # # # |
| # # # #|
|# # # # |
+--------+
`,
		},
		{
			desc:  "Case 2",
			input: "4",
			expected: `+----+
| # #|
|# # |
| # #|
|# # |
+----+
`,
		},
		{
			desc:  "Case 3",
			input: "6x3",
			expected: `+------+
| # # #|
|# # # |
| # # #|
+------+
`,
		},
		{
			desc:  "Case 4",
			input: "1x3",
			expected: `+--+
| #|
|# |
| #|
+--+
`,
		},
		{
			desc:  "Case 5",
			input: "3x1",
			expected: `+---+
| # |
|# #|
+---+
`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str, err := Printer(tC.input)
			assert.Equal(t, tC.expected, str)
			assert.Nil(t, err)
		})
	}
}

func TestErrorsPrinter(t *testing.T) {
	var input string
	var str string
	var err error
	// Incorrect width
	input = "6xText"
	str, err = Printer(input)
	assert.Equal(t, "", str)
	assert.NotNil(t, err)

	// Incorrect height
	input = "Textx6"
	str, err = Printer(input)
	assert.Equal(t, "", str)
	assert.NotNil(t, err)

}
