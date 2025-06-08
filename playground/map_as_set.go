package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return b.Title + " by " + b.Author
}

type bookCollection map[Book]struct{}

func MapAsSet() {
	// Create a new empty book collection
	books := make(bookCollection)

	// Add a book to the collection
	yesMinister := Book{Title: "Yes Minister", Author: "Anthony Jay"}
	books[yesMinister] = struct{}{}

	// Check if the book exists in the collection
	_, ok := books[yesMinister]
	if ok {
		fmt.Println("Book exists")
	} else {
		fmt.Println("Book does not exist")
	}

	// Remove the book from the collection
	delete(books, yesMinister)

	// Iterate through the collection of books
	for book := range books {
		fmt.Println(book)
	}
}
