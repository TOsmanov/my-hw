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
	var inputArr []string

	fmt.Printf("To display a chessboard with equal sides, enter the size of the side.\n")
	fmt.Printf("Example of a square chessboard with side eight: 8\n")
	fmt.Printf("To display a rectangular chessboard, enter the dimensions of the sides in the format: <width>x<height>\n")
	fmt.Printf("Example of a 10 by 6 chessboard: 10x6\n")
	fmt.Scanln(&input)

	inputArr = strings.Split(strings.ToLower(input), "x")

	for i, c := range inputArr {
		var value int
		if i == 0 {
			value, err = strconv.Atoi(c)
			if value >= 2 {
				width = value
			} else {
				width = 2
			}
			if len(inputArr) < 2 {
				height = width
			}
		}
		if i == 1 {
			value, err = strconv.Atoi(c)
			if value >= 2 {
				height = value
			} else {
				height = 2
			}
		}
		if err != nil {
			fmt.Println("Error: you entered an incorrect value")
			panic(err)
		}
	}

	printer.PrintRows(width, height)
}
