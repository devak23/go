package main

import "sort"

// bookCollection is a set of books
type bookCollection map[Book]struct{}

// bookRecommendation is a map whose key is a book, and value is a collection of other books
// Conceptually it looks like this:
// Book A → { Book X, Book Y, Book Z }
// Book B → { Book P, Book Q }
// Book C → { Book M, Book N, Book O }
type bookRecommendation map[Book]bookCollection

type set map[Book]struct{}

func (s set) Contains(book Book) bool {
	_, ok := s[book]
	return ok
}

// newCollection is a helper function to create a new book collection
func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendation)

	// Register all books on everyone's shelf
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	// Recommend a list of related books for each bookworm
	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

// listOtherBooksOnShelves returns a list of books except the one at the given index.
func listOtherBooksOnShelves(bookIndexToRemove int, books []Book) []Book {
	// Initialize the first slice: its capacity is the input slice's length minus 1 and starting length is the number
	// of items up to the index to discard
	otherBooksOnShelves := make([]Book, len(books)-1)
	// Copy the slice of books up to the given index into the initialized index
	copy(otherBooksOnShelves, books[:bookIndexToRemove])
	// Append the rest of the books after the index of the discarded boo into the created slice
	otherBooksOnShelves = append(otherBooksOnShelves, books[bookIndexToRemove+1:]...)
	return otherBooksOnShelves
}

// registerBookRecommendations registers the book recommendations for the given reference book.
// It will create a new book collection if the reference book is not found in the map.
// It will add the other books on the shelves to the book collection for the reference book.
func registerBookRecommendations(recommendations bookRecommendation, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		// check if this reference has already been added to the map
		collection, ok := recommendations[reference]
		if !ok {
			// create a new book collection
			collection = newCollection()
			recommendations[reference] = collection
		}

		// Fill the associated books for the book reference
		collection[book] = struct{}{}
	}
}

// recommendBooks generates a list of book recommendations by comparing a user's shelf with a recommendation dataset.
func recommendBooks(recommendations bookRecommendation, myBooks []Book) []Book {
	bc := make(bookCollection)

	// Register all the books on shelves
	// This step helps us to not recommend a book that has already been read
	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	// Fill recommendations with all the neighboring books
	for _, myBook := range myBooks {
		// Find recommendations in another bookworm's shelves
		for recommendation := range recommendations[myBook] {
			// Find recommendations in another bookworm's shelves
			if !myShelf[recommendation] {
				// Book already on the shelf, so we don't recommend it
				continue
			}

			// Add the book as a recommendation in the collection of books
			bc[recommendation] = struct{}{}
		}
	}

	// Transform the map of books into array
	recommendationsForABook := bookCollectionToListOfBooks(bc)

	return recommendationsForABook
}

// bookCollectionToListOfBooks converts a bookCollection to a sorted slice of Books by author and then by title.
func bookCollectionToListOfBooks(bc bookCollection) []Book {
	books := make([]Book, 0, len(bc))
	for book := range bc {
		books = append(books, book)
	}

	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}
