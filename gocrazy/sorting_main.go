package main

// Goal of this program is to implement sorting using the Go's sort interface
import (
	"fmt"
	"sort"
)

type Flight struct {
	Origin      string
	Destination string
	Price       float64
}

func (f Flight) String() string {
	return fmt.Sprintf("Origin: %s, Destination: %s, Price: %f", f.Origin, f.Destination, f.Price)
}

// ByPrice - defines a type alias that is based on a slice of Flight structs ([]Flight). The purpose of creating this
// type alias is to enable custom sorting behavior. In Go, to use the sort.Sort() function, you need to implement the
// sort.Interface, which requires three methods: Len(), Less(), and Swap(). You can't add methods directly to built-in
// types like []Flight, so by creating the ByPrice type alias, you can then define these required methods on it.
// In this program, the ByPrice type implements sorting by price in ascending order - you can see the Less() method
// on line# 37 compares p[i].Price < p[j].Price. This allows you to convert a regular []Flight slice to ByPrice and
// then sort it by price using sort.Sort(ByPrice(flights)).
type ByPrice []Flight

func (p ByPrice) Len() int {
	return len(p)
}

func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPrice) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

// SortByPrice - sorts the given slice of Flight structs. The sort.Sort(ByPrice(flights)) call modifies the original
// slice directly because slices in Go are reference types - they contain a pointer to the underlying array. When you
// pass a slice to a function, you're passing a copy of the slice header (which includes the pointer), but both the
// original and the copy point to the same underlying array data.
func SortByPrice(flights []Flight) {
	sort.Sort(ByPrice(flights))
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Println(flight)
	}
}

func SortingMain() {
	var flights []Flight

	mumbaiToBangalore := Flight{Origin: "Mumbai", Destination: "Bangalore", Price: 2342}
	mumbaiToChennai := Flight{Origin: "Mumbai", Destination: "Chennai", Price: 4518}
	delhiToChennai := Flight{Origin: "Delhi", Destination: "Chennai", Price: 7566}

	flights = append(flights, mumbaiToChennai)
	flights = append(flights, delhiToChennai)
	flights = append(flights, mumbaiToBangalore)
	fmt.Println("Printing flights before sorting...")
	printFlights(flights)
	fmt.Println("Now sorting...")

	SortByPrice(flights)
	fmt.Println("Printing flights after sorting...")
	printFlights(flights)
}
