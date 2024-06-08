package main

import (
	"fmt"
	"myproject/book"
	"myproject/book_functions"
	"time"
)

func main() {
	books := []book.Book{
		{Name: "Book 1", Author: "Author 1", CreatedAt: time.Now()},
		{Name: "Book 2", Author: "Author 2", CreatedAt: time.Now()},
		{Name: "Book 3", Author: "Author 3", CreatedAt: time.Now()},
		{Name: "Book 4", Author: "Author 4", CreatedAt: time.Now()},
		{Name: "Book 5", Author: "Author 5", CreatedAt: time.Now()},
	}

	bookChannel := make(chan book.Book, len(books))
	filePath := "books.json"

	go func() {
		err := book_functions.WriteBooksToFile(bookChannel, filePath)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}()

	for _, b := range books {
		bookChannel <- b
	}
	close(bookChannel)

	readChannel := make(chan []book.Book)
	go func() {
		err := book_functions.ReadBooksFromFile(readChannel, filePath)
		if err != nil {
			fmt.Println("Error reading from file:", err)
		}
	}()

	readBooks := <-readChannel
	fmt.Println("Books read from file:")
	for _, b := range readBooks {
		fmt.Printf("Name: %s, Author: %s, CreatedAt: %s\n", b.Name, b.Author, b.CreatedAt)
	}
}
