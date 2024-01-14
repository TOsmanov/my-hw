package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

type Comparator struct {
	method int
}

const (
	Size int = iota
	Year
	Rate
)

var id int

func NewBook(title string, author string, year int, size int, rate float32) *Book {
	id++
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func NewComparator(method int) *Comparator {
	return &Comparator{method: method}
}

func (c *Comparator) CompareBooks(book1 *Book, book2 *Book) bool {
	switch c.method {
	case Year:
		return book1.year > book2.year
	case Size:
		return book1.size > book2.size
	case Rate:
		return book1.rate > book2.rate
	default:
		return false
	}
}
