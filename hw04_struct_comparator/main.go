package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw04_struct_comparator/book"
)

func main() {
	// Добавление книг
	books := []*book.Book{
		book.NewBook("978-5-04-154507-9", "Моби Дик", "Герман Мелвилл", 1851, 100, 9.5),
		book.NewBook("978-5-17-109413-3", "Приключения Тома Сойера", "Марк Твен", 1876, 90, 9.7),
		book.NewBook("978-5-04-116639-7", "Приключения Гекльберри Финна", "Марк Твен", 1884, 95, 9.3),
	}

	// Вывод автора книги
	fmt.Println(books[0].Author())

	// Вывод ID книги
	fmt.Printf("Книга «%s» имеет ID: %s\r\n", books[1].Title(), books[1].ID())

	// Вывод книг Марка Твена
	fmt.Println("Книги Марка Твена: ")
	for i := range books {
		if books[i].Author() == "Марк Твен" {
			fmt.Printf("«%s» написана в %d году.\r\n", books[i].Title(), books[i].Year())
		}
	}

	// Изменения полей структуры
	fmt.Printf("Количество страниц книги «%s»: %d.\r\n", books[2].Title(), books[2].Size())
	books[2].SetSize(195)
	fmt.Printf("Количество страниц книги «%s»: %d.\r\n", books[2].Title(), books[2].Size())

	// Сравнение по году
	ComparatorByYear := book.NewComparator(book.Year)
	fmt.Println(ComparatorByYear.CompareBooks(books[0], books[1]))
	fmt.Println(ComparatorByYear.CompareBooks(books[1], books[0]))
	if ComparatorByYear.CompareBooks(books[2], books[1]) {
		fmt.Printf("Книга «%s» была выпущена позже чем книга «%s».\r\n", books[2].Title(), books[1].Title())
	} else {
		fmt.Printf("Книга «%s» была выпущена раньше чем книга «%s».\r\n", books[2].Title(), books[1].Title())
	}

	// Сравнение по рейтингу
	ComparatorByRate := book.NewComparator(book.Rate)
	fmt.Println(ComparatorByRate.CompareBooks(books[0], books[1]))
	fmt.Println(ComparatorByRate.CompareBooks(books[1], books[0]))
}
