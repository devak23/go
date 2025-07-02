package main

import (
	"fmt"
	"github.com/devak23/go/functional-prgs/utils"
)

// One pattern that emerges out of functions being used as first class citizens is the map dispatcher pattern.
// This is the pattern where a map is used to dispatch function calls based on some key.
func add[T utils.Scalar](a, b T) T {
	return a + b
}

func sub[T utils.Scalar](a, b T) T {
	return a - b
}

func mul[T utils.Scalar](a, b T) T {
	return a * b
}

func div[T utils.Scalar](a, b T) T {
	if b == 0 {
		panic("division by zero!")
	}
	return a / b
}

type calculateFunc[T utils.Scalar] func(T, T) T

// getOperations - is the map of an input the user will provide and the operations we support in our calculator. We have
// bound each input to a specific function call. Go does not support generic global variables. You cannot define a
// generic map at the package level. Instead, define a function that returns the map for a specific type parameter.
func getOperations[T utils.Scalar]() map[string]calculateFunc[T] {
	return map[string]calculateFunc[T]{
		"+": add[T],
		"-": sub[T],
		"*": mul[T],
		"/": div[T],
	}
}

func calculateWithMap[T utils.Scalar](a, b T, op string) T {
	if op, ok := getOperations[T]()[op]; ok {
		return op(a, b)
	} else {
		panic(fmt.Sprintf("unsupported operation: %s", op))
	}
}

func main() {
	a, b := 10, 5
	fmt.Println("Add:", calculateWithMap(a, b, "+"))
	fmt.Println("Sub:", calculateWithMap(a, b, "-"))
	fmt.Println("Mul:", calculateWithMap(a, b, "*"))
	fmt.Println("Div:", calculateWithMap(a, b, "/"))

	// The next line will panic because the operation is not supported
	fmt.Println("Mod:", calculateWithMap(a, b, "%"))
}
