package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b := NewBook("978-5-04-154507-9", "Моби Дик", "Герман Мелвилл", 1851, 100, 9.5)
	expected := struct {
		id     string
		title  string
		author string
		year   int
		size   int
		rate   float32
	}{
		id:     "978-5-04-154507-9",
		title:  "Моби Дик",
		author: "Герман Мелвилл",
		year:   1851,
		size:   100,
		rate:   9.5,
	}
	assert.Equal(t, expected.id, b.id)
	assert.Equal(t, expected.title, b.title)
	assert.Equal(t, expected.author, b.author)
	assert.Equal(t, expected.year, b.year)
	assert.Equal(t, expected.size, b.size)
	assert.Equal(t, expected.rate, b.rate)
}

func TestCompareBooks(t *testing.T) {
	b1 := NewBook("978-5-04-154507-9", "Моби Дик", "Герман Мелвилл", 1851, 100, 9.5)
	b2 := NewBook("978-5-17-109413-3", "Приключения Тома Сойера", "Марк Твен", 1876, 90, 9.7)

	// Size
	comparatorSize := NewComparator(Size)
	expectedSize := struct {
		method int
	}{
		method: 0,
	}
	assert.Equal(t, expectedSize.method, comparatorSize.method)
	assert.Equal(t, true, comparatorSize.CompareBooks(b1, b2))
	assert.Equal(t, false, comparatorSize.CompareBooks(b2, b1))

	// Year
	comparatorYear := NewComparator(Year)
	expectedYear := struct {
		method int
	}{
		method: 1,
	}
	assert.Equal(t, expectedYear.method, comparatorYear.method)
	assert.Equal(t, false, comparatorYear.CompareBooks(b1, b2))
	assert.Equal(t, true, comparatorYear.CompareBooks(b2, b1))

	// Rate
	comparatorRate := NewComparator(Rate)
	expectedRate := struct {
		method int
	}{
		method: 2,
	}
	assert.Equal(t, expectedRate.method, comparatorRate.method)
	assert.Equal(t, false, comparatorRate.CompareBooks(b1, b2))
	assert.Equal(t, true, comparatorRate.CompareBooks(b2, b1))

	// Non existent
	comparatorNone := NewComparator(3)
	expectedNone := struct {
		method int
	}{
		method: 3,
	}
	assert.Equal(t, expectedNone.method, comparatorNone.method)
	assert.Equal(t, false, comparatorNone.CompareBooks(b1, b2))
}

func TestSetAndGet(t *testing.T) {
	b1 := NewBook("", "", "", 0, 0, 0)
	b1.SetID("978-5-04-154507-9")
	assert.Equal(t, "978-5-04-154507-9", b1.ID())
	b1.SetTitle("Моби Дик")
	assert.Equal(t, "Моби Дик", b1.Title())
	b1.SetAuthor("Герман Мелвилл")
	assert.Equal(t, "Герман Мелвилл", b1.Author())
	b1.SetYear(1851)
	assert.Equal(t, 1851, b1.Year())
	b1.SetSize(100)
	assert.Equal(t, 100, b1.Size())
	b1.SetRate(9.5)
	assert.Equal(t, float32(9.5), b1.Rate())

}
