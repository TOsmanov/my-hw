package printer

import (
	"fmt"
	"strings"
)

func printLine(even bool, width int) {
	var a, b = "#", " "
	if even {
		a, b = b, a
	}
	for i := 0; i < width; i++ {
		if i%2 == 0 {
			fmt.Print(a)
		} else {
			fmt.Print(b)
		}
	}
}

func PrintRows(width int, height int) {
	var border = strings.Repeat("-", width)
	fmt.Printf("+%s+\n", border)
	for i := 0; i < height; i++ {
		fmt.Print("|")
		if i%2 == 0 {
			printLine(true, width)
		} else {
			printLine(false, width)
		}
		fmt.Print("|\n")
	}
	fmt.Printf("+%s+\n", border)
}
