package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworm's shelf
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// A Book describes a book on a bookworm's shelf
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b Book) String() string {
	return fmt.Sprintf("%s: %s", b.Title, b.Author)
}

// To sort the books in the output of the main function, we used sort.Slice, which takes a function as the sorting
// strategy. There is a second option the sort package provides sort.Interface, which can be implemented to sort slices
// or user-defined collections. It becomes very handy when implementing custom sorting — in our case, by author and title.
// This sorting applies only to collection so we add an intermediate custom type representing the collection of Book. The
// following listing shows the implementation of sort.Interface with a type called byAuthorThenTitle.
type byAuthorThenTitle []Book

// Len implements sort.Interface by returning the length of the BooksByAuthor slice.
func (b byAuthorThenTitle) Len() int { return len(b) }

// Swap implements sort.Interface and swaps 2 books
func (b byAuthorThenTitle) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements sort.Interface and returns books sorted by Author first and then by Title
func (b byAuthorThenTitle) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// Each Go field is tagged with the name of the JSON field. Note that the name of the field doesn’t have to match the
// name of the tag - it’s simply a convention that makes it more readable. Another convention is that fields that are
// slices should be named with plural word.

// loadBookworms reads the file and returns a list of bookworms and their beloved books
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err.Error())
			return
		}
	}(f)

	// Decode the file and store the content int the value bookworms
	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bookworms, nil
}

// booksCount registers all the books and their respective occurrences from the bookworm shelves
func booksCount(bookworms []Bookworm) map[Book]uint {
	countByBook := make(map[Book]uint) // create a map {book, count} where key is book and # of books as count
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			countByBook[book]++
		}
	}
	return countByBook
}

// findCommonBooks returns books that are on more than one bookworm's shelf
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)
	// Now that we’ve counted the number of copies of each book on every bookshelf, the next step is to loop over all
	// of them and keep those with more than one copy. Let’s declare a slice that will contain all the books that were
	// found multiple times in the collections of all bookworms
	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

func sortBooks(books []Book) []Book {
	sort.Sort(byAuthorThenTitle(books))
	return books
}

// sortBooks sorts the books by Author first and then title
func sortBooksAlt(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})

	return books
}

// displayBooks displays the title and Author of each book on the console
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
