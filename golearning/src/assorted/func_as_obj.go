package assorted

import "fmt"

// define an "interface" (or a type - in GO) called testInt which is any
// function that accepts an integer and returns a boolean
type testInt func(int) bool

// isEven is an even function that will determine if a number is even.
// this function follows the pattern of the type defined above
func isEven(a int) bool {
	return a%2 == 0
}

// isEven is an even function that will determine if a number is odd.
// this function follows the pattern of the type defined above
func isOdd(b int) bool {
	return b%2 == 1
}

// now define a function that accepts function as an argument
func filter(slice []int, fn testInt) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// FuncAsObjMain is the main function that's invoked from main.go
func FuncAsObjMain() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers := filter(slice, isEven)
	oddNumbers := filter(slice, isOdd)

	fmt.Println("even numbers =", evenNumbers)
	fmt.Println("odd numbers = ", oddNumbers)
}
