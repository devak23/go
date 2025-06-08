package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return b.Title + " by " + b.Author
}

// setOfBooks defines a new type named that is based on a map. This is a common go idiom for implementing a set (which
// go doesn't have a built-in type). The empty struct takes zero memory, making it ideal for sets where you only care
// about the presence or absence of an element and not the associated value. You can see that the setOfBooks uses Book
// struct as the key. This means each book can appear only once in the set. Using struct{} as a value-type (instead of a
// boolean or an integer) is memory-efficient since an empty struct occupies no memory.
type setOfBooks map[Book]struct{}

func UsingMapAsSet() {
	// Create a new empty book collection
	books := make(setOfBooks)

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
