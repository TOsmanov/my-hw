package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TOsmanov/my-hw/hw03_chessboard/printer"
)

func main() {
	var input string
	var err error
	var width int
	var height int

	fmt.Println("To display a chessboard with equal sides, enter the size of the side.")
	fmt.Println("Example of a square chessboard with side eight: 8")
	fmt.Println("To display a rectangular chessboard, enter the dimensions of the sides in the format: <width>x<height>")
	fmt.Println("Example of a 10 by 6 chessboard: 10x6")
	fmt.Scanln(&input)

	if len(input) == 0 {
		printer.PrintRows(8, 8)
	}

	inputArr := strings.Split(strings.ToLower(input), "x")

	width, err = strconv.Atoi(inputArr[0])
	if width < 2 {
		width = 2
	}
	if len(inputArr) < 2 {
		height = width
	} else {
		height, err = strconv.Atoi(inputArr[1])
		if height < 2 {
			height = 2
		}
	}
	if err != nil {
		fmt.Printf("Error: the entered size \"%s\" is incorrect.\n", input)
		panic(err)
	}

	printer.PrintRows(width, height)
}
