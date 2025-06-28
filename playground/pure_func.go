package main

import (
	"os"
	"strconv"
)

func main() {
	// strconv.Atoi converts a string to an integer.
	// os.Args[0] is the program name, so we start from os.Args[1] for the first argument.
	//x, _ := strconv.Atoi(os.Args[1])
	//y, _ := strconv.Atoi(os.Args[2])
	//os.Exit(addScalar(x, y))

	// ----------- OR --------------

	// os.Exit(add(arg(0), arg(1)))

	// ----------- OR --------------
	//var sum int
	//for _, v := range os.Args[1:] {
	//	x, _ := strconv.Atoi(v)
	//	sum = add(sum, x)
	//}
	//os.Exit(sum)

	// ----------- OR --------------

	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		accumulate(x)
	}
	os.Exit(y)
}

var y int // y is a package-level variable that accumulates the sum of integers passed as command-line arguments.

// accumulate adds the given integer x to the package-level variable y.
func accumulate(x int) {
	y += x
}

func arg(n int) (r int) {
	r, _ = strconv.Atoi(os.Args[n+1])
	return
}

func add(a, b int) int {
	return a + b
}

// Scalar is a type constraint that allows any numeric type. The ~ operator allows types that have the specified
// underlying type (e.g., custom types based on int). The | operator is used to combine multiple types. This constraint
// enables generic functions to accept any integer or floating-point type, including user-defined types with those
// underlying types.
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func addScalar[T Scalar](x, y T) T {
	return x + y
}
