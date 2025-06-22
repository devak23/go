package gobyexample

import "fmt"
import "math"

// const declares a constant value
const s string = "constant"

func ConstantsMain() {
	fmt.Println(s)
	const n = 500000000
	// const expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until it's given one such by an
	// explicit cast
	fmt.Println(int64(d))

	// a number can be given a type by using it in a context, for
	// instance a variable assignment or a function call.
	// Here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
