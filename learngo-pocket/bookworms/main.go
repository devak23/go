package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: Failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(bookworms)
}
