package book_functions

import (
	"bufio"
	"encoding/json"
	"myproject/book"
	"os"
)

func WriteBooksToFile(ch <-chan book.Book, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for b := range ch {
		data, err := json.Marshal(b)
		if err != nil {
			return err
		}
		writer.WriteString(string(data) + "\n")
	}
	return nil
}

func ReadBooksFromFile(ch chan<- []book.Book, filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var books []book.Book
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var b book.Book
		err := json.Unmarshal(scanner.Bytes(), &b)
		if err != nil {
			return err
		}
		books = append(books, b)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	ch <- books
	close(ch)
	return nil
}
