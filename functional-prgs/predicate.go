package main

import . "fmt"

type predicate func(int) bool

func filter(slice []int, condition predicate) []int {
	var result []int

	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	// Example usage of filter function with a predicate
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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
