package assorted

import "fmt"

// important thing to learn from this example is that int and float numeric
// types cannot be mixed as is. they need to be converted into appropriate data
// types before performing the operations. This is a compile time check

func intAndFloatMain() {
	var k float64 = 3.14159
	var m int = 2
	var result1 = int(k) + m

	var result2 int = int(k) * m
	fmt.Printf("%T %T\n", result1, result2)
	fmt.Printf("%v %v\n", result1, result2)

	var result3 float64 = k + float64(m)
	var result4 float64 = k * float64(m)

	fmt.Printf("%F %F\n", result3, result4)
	fmt.Printf("%v %v\n", result3, result4)
}
