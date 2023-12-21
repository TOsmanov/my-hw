package printer

import (
	"fmt"
	"strings"
)

func printLine(even bool, size int) {
	var a, b = "#", " "
	if even {
		a, b = b, a
	}
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			fmt.Print(a)
		} else {
			fmt.Print(b)
		}
	}
}

func PrintRows(size int) {
	var border = strings.Repeat("-", size)
	fmt.Printf("+%s+\n", border)
	for i := 0; i < size; i++ {
		fmt.Print("|")
		if i%2 == 0 {
			printLine(true, size)
		} else {
			printLine(false, size)
		}
		fmt.Print("|\n")
	}
	fmt.Printf("+%s+\n", border)
}
