package book

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

type bookJSON struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float32 `json:"rate,omitempty"`
}

const (
	maxBookSize int = 1000
)

func (b Book) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(&bookJSON{b.ID, b.Title, b.Author, b.Year, b.Size, b.Rate})
	return j, err
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var book bookJSON
	err := json.Unmarshal(data, &book)
	b.ID = book.ID
	b.Title = book.Title
	b.Author = book.Author
	b.Year = book.Year
	b.Size = book.Size
	b.Rate = book.Rate
	return err
}

func MarshalSlice(books []Book) ([]byte, error) {
	size := maxBookSize * len(books)
	arr := make([]byte, 1, size)
	arr[0] = byte('[')
	for i, b := range books {
		j, err := b.MarshalJSON()
		if err != nil {
			return nil, err
		}
		arr = append(arr, j...)
		if i != len(books)-1 {
			arr = append(arr, byte(','))
		}
	}
	arr = append(arr, byte(']'))
	return arr, nil
}

func UnmarshalSlice(data []byte) ([]Book, error) {
	booksJ := strings.Split(string(data), "},{")
	fmt.Println(booksJ)
	l := len(booksJ)
	books := make([]Book, l)
	for i, b := range booksJ {
		var builder strings.Builder
		switch i {
		case 0:
			builder.WriteString(b[1:])
			builder.WriteRune('}')
		case l - 1:
			builder.WriteRune('{')
			builder.WriteString(b[:len(b)-1])
		default:
			builder.WriteRune('{')
			builder.WriteString(b)
			builder.WriteRune('}')
		}
		book := []byte(builder.String())
		err := books[i].UnmarshalJSON(book)
		if err != nil {
			return nil, err
		}
	}
	return books, nil
}
