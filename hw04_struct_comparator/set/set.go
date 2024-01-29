package set

import "github.com/TOsmanov/my-hw/hw04_struct_comparator/book"

func (b *book.Book) SetID(id string) {
	b.ID = id
}

func (b *book.Book) SetTitle(title string) {
	b.Title = title
}

func (b *book.Book) SetAuthor(author string) {
	b.Author = author
}

func (b *book.Book) SetYear(year int) {
	b.Year = year
}

func (b *book.Book) SetSize(size int) {
	b.Size = size
}

func (b *book.Book) SetRate(rate float32) {
	b.Rate = rate
}
