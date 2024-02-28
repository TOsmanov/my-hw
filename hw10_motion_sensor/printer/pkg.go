package printer

import (
	"fmt"
)

func Printer(inputCh chan int) {
	for {
		output := <-inputCh
		fmt.Println(output)
	}
}
