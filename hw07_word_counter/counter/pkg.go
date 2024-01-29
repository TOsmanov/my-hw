package counter

import (
	"strings"
)

func CountWords(text string) map[string]int {
	chars := []rune(text)
	word := make([]rune, 0, 40)
	output := make(map[string]int)
	for i, char := range chars {
		switch {
		case i == len(chars)-1:
			switch char {
			case 9, 10, 32, 33, 42, 44, 45, 46, 58, 59:
			default:
				word = append(word, char)
			}
			if len(word) >= 1 {
				w := strings.ToLower(string(word))
				val, ok := output[w]
				if ok {
					output[w] = val + 1
				} else {
					output[w] = 1
				}
			}
		case char == 9, char == 10, char == 32:
			if len(word) >= 1 {
				w := strings.ToLower(string(word))
				val, ok := output[w]
				if ok {
					output[w] = val + 1
				} else {
					output[w] = 1
				}
			}
			word = make([]rune, 0, 40)
		case char == 58, char == 59, char == 42, char == 44:
		case char == 45, char == 46, char == 33, char == 63:
		default:
			word = append(word, char)
		}
	}
	return output
}
