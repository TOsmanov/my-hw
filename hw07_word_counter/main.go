package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw07_word_counter/counter"
)

func main() {
	text := `Hello world!`

	fmt.Println(counter.CountWords(text))
}
