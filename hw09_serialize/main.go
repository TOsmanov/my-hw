package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/TOsmanov/my-hw/hw09_serialize/book"
	pbook "github.com/TOsmanov/my-hw/hw09_serialize/pbook"
	"google.golang.org/protobuf/proto"
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

	j, err := book.Book.MarshalJSON(b1)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

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

	j, err = json.Marshal(books)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Marshal slice:\n\t %v\n", string(j))

	booksCopy := []book.Book{}
	json.Unmarshal(j, &booksCopy)
	fmt.Printf("Unmarshal slice:\n\t %v\n\n", booksCopy)

	// Slices methods

	j, err = book.MarshalSlice(books)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Marshal slice method:\n\t %v\n", string(j))

	books1Copy, err := book.UnmarshalSlice(j)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Unmarshal slice method:\n\t %v\n\n", books1Copy)

	// Proto

	pBook1 := pbook.Book{
		Id:     "978-5-04-154507-9",
		Title:  "Моби Дик",
		Author: "Герман Мелвилл",
		Year:   1851,
		Size:   576,
		Rate:   9,
	}
	j, err = proto.Marshal(pBook1.ProtoReflect().Interface())
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Proto Marshal:\n\t %v\n", strings.Trim(string(j), "\n"))
	var pBookCopy pbook.Book
	proto.Unmarshal(j, &pBookCopy)
	fmt.Printf("Proto Unmarshal:\n\t %v\n\n", pBookCopy.String())

	// Proto slices

	pBooks := pbook.Books{
		Book: []*pbook.Book{
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
		},
	}

	j, err = proto.Marshal(pBooks.ProtoReflect().Interface())
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Proto Marshal slice:\n\t %v\n", strings.Trim(string(j), "\n"))

	var pBooksCopy pbook.Books
	proto.Unmarshal(j, &pBooksCopy)
	fmt.Printf("Proto Unmarshal slice:\n\t %v\n\n", pBooksCopy.String())
}
