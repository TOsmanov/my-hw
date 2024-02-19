package book

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	books = []Book{
		{
			ID:     "978-5-04-154507-9",
			Title:  "Моби Дик",
			Author: "Герман Мелвилл",
			Year:   1851,
			Size:   576,
			Rate:   9,
		},
		{
			ID:     "978-5-17-109413-3",
			Title:  "Приключения Тома Сойера",
			Author: "Марк Твен",
			Year:   1876,
			Size:   90,
		},
		{
			ID:     "978-5-04-116639-7",
			Title:  "Приключения Гекльберри Финна",
			Author: "Марк Твен",
			Year:   1884,
			Size:   95,
			Rate:   9.3,
		},
	}

	book1 = Book{
		ID:     "978-5-17-109413-3",
		Title:  "Приключения Тома Сойера",
		Author: "Марк Твен",
		Year:   1876,
		Size:   90,
		Rate:   9.7,
	}

	book2 = Book{
		ID:     "978-5-04-116639-7",
		Title:  "Приключения Гекльберри Финна",
		Author: "Марк Твен",
		Year:   1884,
		Size:   95,
		Rate:   9.3,
	}
)

func TestMarshal(t *testing.T) {
	b1 := book1
	expected := `{"id":"978-5-17-109413-3","title":"Приключения Тома Сойера","author":"Марк Твен",
	"year":1876,"size":90,"rate":9.7}`
	expected = strings.Join(strings.Split(expected, "\n\t"), "")
	res, err := json.Marshal(b1)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(res))
}

func TestUnmarshal(t *testing.T) {
	data := `{"id":"978-5-04-116639-7","Title":"Приключения Гекльберри Финна","Author":"Марк Твен",
	"Year":1884,"Size":95,"Rate":9.3}`
	data = strings.Join(strings.Split(data, "\n\t"), "")
	expected := book2
	var book Book
	err := json.Unmarshal([]byte(data), &book)
	assert.Nil(t, err)
	assert.Equal(t, expected, book)
}

func TestMarshalSlice(t *testing.T) {
	books1 := books
	expected := `[{"id":"978-5-04-154507-9","title":"Моби Дик","author":"Герман Мелвилл","year":1851,
	"size":576,"rate":9},{"id":"978-5-17-109413-3","title":"Приключения Тома Сойера","author":"Марк Твен",
	"year":1876,"size":90},{"id":"978-5-04-116639-7","title":"Приключения Гекльберри Финна",
	"author":"Марк Твен","year":1884,"size":95,"rate":9.3}]`
	expected = strings.Join(strings.Split(expected, "\n\t"), "")

	data, err := MarshalSlice(books1)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(data))
}

func TestUnmarshalSlice(t *testing.T) {
	data := `[{"id":"978-5-04-154507-9","title":"Моби Дик","author":"Герман Мелвилл","year":1851,
	"size":576,"rate":9},{"id":"978-5-17-109413-3","title":"Приключения Тома Сойера","author":"Марк Твен",
	"year":1876,"size":90},{"id":"978-5-04-116639-7","title":"Приключения Гекльберри Финна",
	"author":"Марк Твен","year":1884,"size":95,"rate":9.3}]`
	data = strings.Join(strings.Split(data, "\n\t"), "")

	expected := books
	b, err := UnmarshalSlice([]byte(data))
	assert.Nil(t, err)
	assert.Equal(t, expected, b)
}
