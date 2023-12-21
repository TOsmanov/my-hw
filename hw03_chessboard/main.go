package main

import (
	"fmt"
	"strconv"

	"github.com/TOsmanov/my-hw/hw03_chessboard/printer"
)

func main() {
	var input string
	var err error
	var size int

	fmt.Printf("Enter the size of the chessboard (min: 2): ")
	fmt.Scanln(&input)

	size, err = strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: you entered an incorrect value")
		panic(err)
	}

	if size < 2 {
		size = 2
	}

	printer.PrintRows(size)
}
