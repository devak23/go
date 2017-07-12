package main

import "fmt"

// important thing to learn from this example is that int and float numeric
// types cannot be mixed as is. they need to be converted into appropriate data
// types before performing the operations. This is a compile time check

func main() {
	var k float64 = 3.14159
	var m int = 2
	var result1 = int(k) + m
	fmt.Println("result1 = ", result1)

	var result2 int = int(k) * m
	fmt.Println("result2 = ", result2)

	var result3 float64 = k + float64(m)
	fmt.Println("result3 = ", result3)

	var result4 float64 = k * float64(m)
	fmt.Println("result4 = ", result4)
}