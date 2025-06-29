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
