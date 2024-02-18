package main

import (
	"encoding/json"
	"fmt"

	"github.com/TOsmanov/my-hw/hw09_serialize/book"
	"github.com/TOsmanov/my-hw/hw09_serialize/proto"
)

func main() {
	b1 := book.Book{
		ID:     "978-5-04-154507-9",
		Title:  "Моби Дик",
		Author: "Герман Мелвилл",
		Year:   1851,
		Size:   576,
		Rate:   9,
	}

	j, _ := book.Book.MarshalJSON(b1)
	fmt.Printf("Marshal:\n\t %v\n", string(j))

	var bookCopy book.Book

	json.Unmarshal(j, &bookCopy)
	fmt.Printf("Unmarshal:\n\t %v\n\n", bookCopy)

	// Slices

	books := []book.Book{
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
			Rate:   9.7,
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

	j, _ = json.Marshal(books)
	fmt.Printf("Marshal slice:\n\t %v\n", string(j))

	booksCopy := []book.Book{}
	json.Unmarshal(j, &booksCopy)
	fmt.Printf("Unmarshal slice:\n\t %v\n\n", booksCopy)

	// Proto

	book1 := proto.Book{
		Id:     "978-5-04-154507-9",
		Title:  "Моби Дик",
		Author: "Герман Мелвилл",
		Year:   1851,
		Size:   576,
		Rate:   9,
	}
	fmt.Printf("Proto String:\n\t %v\n\n", book1.String())
	fmt.Printf("Proto String:\n\t %v\n\n", book1.GetTitle())

	pBooks := []proto.Book{
		{
			Id:     "978-5-04-154507-9",
			Title:  "Моби Дик",
			Author: "Герман Мелвилл",
			Year:   1851,
			Size:   576,
			Rate:   9,
		},
		{
			Id:     "978-5-17-109413-3",
			Title:  "Приключения Тома Сойера",
			Author: "Марк Твен",
			Year:   1876,
			Size:   90,
			Rate:   9.7,
		},
		{
			Id:     "978-5-04-116639-7",
			Title:  "Приключения Гекльберри Финна",
			Author: "Марк Твен",
			Year:   1884,
			Size:   95,
			Rate:   9.3,
		},
	}
	fmt.Printf("Proto String:\n\t %v\n\n", pBooks[0].String())
	fmt.Printf("Proto String:\n\t %v\n\n", pBooks[0].GetTitle())
	fmt.Printf("Proto String:\n\t %v\n\n", pBooks[1].GetTitle())
}
