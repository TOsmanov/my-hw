package get

import (
	"github.com/TOsmanov/my-hw/hw04_struct_comparator/book"
)

type MyBook struct {
	*book.Book
}

func (b *MyBook) ID() string {
	return b.ID
}

// func (b *Book) Title() string {
// 	return b.Title
// }

// func (b *book.Book) Author() string {
// 	return b.Author
// }

// func (b *book.Book) Year() int {
// 	return b.Year
// }

// func (b *book.Book) Size() int {
// 	return b.Size
// }

// func (b *book.Book) Rate() float32 {
// 	return b.Rate
// }
