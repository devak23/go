package main

import (
	"encoding/json"
	"fmt"
	"os"
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
