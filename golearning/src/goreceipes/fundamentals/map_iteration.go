package fundamentals

import (
	"fmt"
	"sort"
)

// MapIterationMain is the main function to be invoked from main.go
func MapIterationMain() {
	// Initialize map with key as int and string as value
	scientists := make(map[int]string)

	// Add the data as key value pairs
	scientists[1] = "Sir Issac Newton"
	scientists[2] = "Albert Einsten"
	scientists[0] = "Neils Bohr"
	scientists[4] = "Paul Dirac"
	scientists[3] = "Subramanyan Chandrashekhar"

	// Define a slice (keys) for specifying the order of the map
	var keys []int

	// Append keys of the map
	for k := range scientists {
		keys = append(keys, k)
	}

	// Use the sort.Ints() to sort a slice of ints in increasing order
	sort.Ints(keys)

	// Now iterate the map with order
	for k := range keys {
		fmt.Printf("%d =  %s\n", k, scientists[k])
	}
}
