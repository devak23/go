package gobyexample

import "fmt"

func VariablesMain() {
	// var declares one or more variables
	var aString = "initial"
	fmt.Println(aString)

	// multiple variables can be declared in a single line
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go will infer the type of variables
	var d = true
	fmt.Println(d)

	// variables without initialization are defaulted to 0
	var e int
	fmt.Println(e)

	// := shorthand for declaring and initializing
	f := "short"
	fmt.Println(f)
}
