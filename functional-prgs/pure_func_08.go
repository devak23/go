package main

import (
	"github.com/devak23/go/functional-prgs/utils"
	"os"
	"strconv"
)

// ScalarAccumulator - defines an accumulator function type that takes a scalar value of type T and returns a
// function that takes a scalar value and returns the same type T.
type ScalarAccumulator[T utils.Scalar] func(T) T

// MakeAccumulator - is the factory function that creates an accumulator for a scalar type T.
func MakeAccumulator[T utils.Scalar]() ScalarAccumulator[T] {
	var y T // Global variable that will keep getting updated with the sum of integers passed. This variable is captured
	// by the closure and persists between function calls
	return func(x T) T {
		y += x
		return y
	}
}

func main() {
	a := MakeAccumulator[int]() // Create an accumulator function for integers

	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x) // Call the accumulator function with the integer value x. This calls the closure that updates the global variable y
	}

	os.Exit(a(0)) // Pass 0 to get the accumulated value, as the accumulator function returns the sum.
}
