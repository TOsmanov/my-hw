package printer

import (
	"fmt"
	"strconv"
	"strings"
)

func printLine(even bool, width int) string {
	a := '#'
	b := ' '
	var str string
	if even {
		a, b = b, a
	}
	for i := 0; i < width; i++ {
		if i%2 == 0 {
			str += string(a)
		} else {
			str += string(b)
		}
	}
	return str
}

func PrintRows(width int, height int) string {
	var str string

	border := strings.Repeat("-", width)
	str += fmt.Sprintf("+%s+\n", border)
	for i := 0; i < height; i++ {
		str += "|"
		if i%2 == 0 {
			str += printLine(true, width)
		} else {
			str += printLine(false, width)
		}
		str += fmt.Sprintln("|")
	}
	str += fmt.Sprintf("+%s+\n", border)
	return str
}

func Printer(input string) (string, error) {
	var err error
	var width int
	var height int

	if len(input) == 0 {
		return PrintRows(8, 8), nil
	}

	inputArr := strings.Split(strings.ToLower(input), "x")

	width, err = strconv.Atoi(inputArr[0])
	if err != nil {
		return "", fmt.Errorf("error incorrect size: the entered width \"%s\" is incorrect", input)
	} else if width < 2 {
		width = 2
	}
	if len(inputArr) < 2 {
		height = width
	} else {
		height, err = strconv.Atoi(inputArr[1])
		if err != nil {
			return "", fmt.Errorf("error incorrect size: the entered height \"%s\" is incorrect", input)
		} else if height < 2 {
			height = 2
		}
	}
	return PrintRows(width, height), err
}
