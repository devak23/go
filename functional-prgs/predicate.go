package main

import (
	. "fmt"
	"github.com/devak23/go/functional-prgs/utils"
)

type predicate[T utils.Scalar] func(T) bool

func filter[T utils.Scalar](slice []T, condition predicate[T]) []T {
	var result []T

	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	numbers := []int{11, 2, 3, 41, 5, 6, 7, 18, 9, 10}
	evenNumbers := filter(numbers, isEven)
	Printf("Even numbers: %v", evenNumbers)

	largerThanFive := filter(numbers, isLargerThanFive)
	Printf("\nNumbers larger than five: %v", largerThanFive)

	oddNumbers := filter(numbers, isOdd)
	Printf("\nOdd numbers: %v\n", oddNumbers)

	smallerThanFive := func(i int) bool { return i < 5 } // inline function
	Printf("Numbers smaller than five: %v\n", filter(numbers, smallerThanFive))

	Printf("Numbers greater than 20: %v\n", filter(numbers, func(i int) bool { return i > 20 }))

	Printf("Numbers smaller than 20: %v\n", filter(numbers, createSmallerThanPredicate(20)))

	// --------------------- Slices of predicates ---------------------
	Println("\n\nUsing named predicates: ----------------------")
	namedPredicates := []namedPredicate[int]{
		{"isEven", isEven},
		{"isLargerThanFive", isLargerThanFive},
		{"isOdd", isOdd},
	}
	for _, np := range namedPredicates {
		Printf("Numbers that satisfy predicate %s: %v\n", np.name, filter(numbers, np.fn))
	}
}

type namedPredicate[T utils.Scalar] struct {
	name string
	fn   predicate[T]
}

// createLargerThanPredicate is a function that creates a predicate.
func createSmallerThanPredicate[T utils.Scalar](threshold T) predicate[T] {
	return func(n T) bool {
		return n < threshold
	}
}

func isEven(n int) bool {
	return n%2 == 0
}

func isLargerThanFive(n int) bool {
	return n > 5
}

func isOdd(n int) bool {
	return n%2 != 0
}
