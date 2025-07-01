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

	cc := ConstraintChecker[int]{
		largerThan:  createLargerThanPredicate(5),
		smallerThan: createSmallerThanPredicate(20),
	}
	Printf("\nConstraint checker: %v", cc.check(10))
	result1 := cc.check(10) // true, because 10 > 5 and 10 < 20
	result2 := cc.check(4)  // false, because 4 is not > 5
	result3 := cc.check(25) // false, because 25 is not < 20

	Printf("\nresult1: %v, result2: %v, result3: %v\n", result1, result2, result3)
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

func createLargerThanPredicate[T utils.Scalar](threshold T) predicate[T] {
	return func(t T) bool {
		return t > threshold
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

type ConstraintChecker[T utils.Scalar] struct {
	largerThan  predicate[T]
	smallerThan predicate[T]
}

func (cc *ConstraintChecker[T]) check(t T) bool {
	return cc.largerThan(t) && cc.smallerThan(t)
}
