package counter

import (
	"strings"
)

func CountWords(text string) map[string]int {
	chars := []rune(text)
	word := make([]rune, 0, 40)
	output := make(map[string]int)
	for i, char := range chars {
		if i != len(chars)-1 {
			switch char {
			case 32, 9, 10:
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
			case 58, 59, 42, 44, 45, 46, 33, 63:
			default:
				word = append(word, char)
			}
		} else {
			switch char {
			case 32, 9, 10, 33, 42, 44, 45, 46, 58, 59, 63:
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
		}
	}
	return output
}
