package main

import "fmt"

// defines a complex data structure that represents a book recommendation system using nested maps.
// Each book (key) has a collection of other books (value) associated with it
// This represents a "if you like this book, you might also like these books" recommendation system
type bookRecommendation map[Book]setOfBooks

// This data structure is particularly efficient for:
// 1. Quickly finding all books recommended for a given book
// 2. Checking if a specific book is recommended for another book
// 3. Ensuring no duplicate recommendations (since each inner map is a set)
//It's a common pattern in recommendation systems where relationships between items need to be tracked efficiently.

func UsingRecommendationSystem() {
	// Create a new recommendation system - an empty map/set.
	recommendations := make(bookRecommendation)

	// For specific books, create a collection of recommended books
	theHobbit := Book{Title: "The Hobbit", Author: "J.R.R Tolkien"} // <-- Specific book
	recommendations[theHobbit] = make(setOfBooks)

	// Create the recommended books for theHobbit. This will generally be via some "intelligent algorithm", but we will
	// hard code here for simplicity.
	aBriefHistoryOfTime := Book{Title: "A Brief History of Time", Author: "Stephen Hawking"}
	sherlockHolmes := Book{Title: "Sherlock Holmes", Author: "Sir Arthur Conan Doyle"}

	// Now we add these books to the recommendation system (map/set) as recommended books for the person reading "The Hobbit"
	recommendations[theHobbit][aBriefHistoryOfTime] = struct{}{}
	recommendations[theHobbit][sherlockHolmes] = struct{}{}

	// Note that the struct{}{} is an empty struct that is used as a marker value. It literally takes no memory.

	// Check if a specific book is recommended for theHobbit
	_, ok := recommendations[theHobbit][sherlockHolmes]
	if ok {
		fmt.Println(sherlockHolmes.Title, "is recommended for the reader of", theHobbit.Title)
	} else {
		fmt.Println(sherlockHolmes.Title, "is not recommended for the reader of", theHobbit.Title)
	}

	// Iterate through the collection of recommendations for theHobbit
	fmt.Println("If you enjoyed reading", theHobbit.Title, "you might also like these books:")
	for recommendedBook := range recommendations[theHobbit] {
		fmt.Println("-", recommendedBook.Title, "by", recommendedBook.Author)
	}
}
