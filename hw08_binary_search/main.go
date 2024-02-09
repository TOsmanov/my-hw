package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw08_binary_search/search"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, found := search.BinarySearch(list, 4)
	if found {
		fmt.Printf("Found: %d\n", result)
	} else {
		fmt.Println("Nothing found!")
	}
}
