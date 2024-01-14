package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw04_struct_comparator/book"
)

func main() {
	// Добавление книг
	books := []*book.Book{
		book.NewBook("Моби Дик", "Герман Мелвилл", 1851, 100, 9.5),
		book.NewBook("Приключения Тома Сойера", "Марк Твен", 1876, 90, 9.7),
		book.NewBook("Приключения Гекльберри Финна", "Марк Твен", 1884, 95, 9.3),
	}

	// Вывод автора книги
	fmt.Println(books[0].GetAuthor())

	// Вывод книг Марка Твена
	fmt.Println("Книги Марка Твена: ")
	for i := range books {
		if books[i].GetAuthor() == "Марк Твен" {
			fmt.Printf("«%s» написана в %d году.\r\n", books[i].GetTitle(), books[i].GetYear())
		}
	}

	// Изменения полей структуры
	fmt.Printf("Количество страниц книги «%s»: %d.\r\n", books[2].GetTitle(), books[2].GetSize())
	books[2].SetSize(195)
	fmt.Printf("Количество страниц книги «%s»: %d.\r\n", books[2].GetTitle(), books[2].GetSize())

	// Сравнение по году
	ComparatorByYear := book.NewComparator(book.Year)
	fmt.Println(ComparatorByYear.CompareBooks(books[0], books[1]))
	fmt.Println(ComparatorByYear.CompareBooks(books[1], books[0]))
	if ComparatorByYear.CompareBooks(books[2], books[1]) {
		fmt.Printf("Книга «%s» была выпущена позже чем книга «%s».\r\n", books[2].GetTitle(), books[1].GetTitle())
	} else {
		fmt.Printf("Книга «%s» была выпущена раньше чем книга «%s».\r\n", books[2].GetTitle(), books[1].GetTitle())
	}

	// Сравнение по рейтингу
	ComparatorByRate := book.NewComparator(book.Rate)
	fmt.Println(ComparatorByRate.CompareBooks(books[0], books[1]))
	fmt.Println(ComparatorByRate.CompareBooks(books[1], books[0]))
}
