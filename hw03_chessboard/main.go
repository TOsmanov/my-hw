package main

import (
	"fmt"
	"os"

	"github.com/TOsmanov/my-hw/hw03_chessboard/printer"
)

func main() {
	var input string

	fmt.Println("To display a chessboard with equal sides, enter the size of the side.")
	fmt.Println("Example of a square chessboard with side eight: 8")
	fmt.Println("To display a rectangular chessboard, enter the dimensions of the sides in the format: <width>x<height>")
	fmt.Println("Example of a 10 by 6 chessboard: 10x6")
	fmt.Scanln(&input)

	str, err := printer.Printer(input)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(str)
}
