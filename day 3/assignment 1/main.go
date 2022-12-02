package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// type Book struct {
// 	title, author, genre, publisher string
// }

func main() {
	f, err := os.Open("books.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	records = records[1:]
	// books := []Book{}

	for _, r := range records {
		// books = append(books, Book{title: r[0], author: r[1], genre: r[2], publisher: r[3]})
		row := "Title: " + r[0] + "\nAuthor: " + r[1] + "\nGenre: " + r[2] + "\nPublisher: " + r[3] + "\n"
		fmt.Println(row)
	}

	// fmt.Println(books)
}
