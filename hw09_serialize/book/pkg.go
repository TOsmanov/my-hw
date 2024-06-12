package book

import (
	"encoding/json"
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
	j, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func UnmarshalSlice(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
