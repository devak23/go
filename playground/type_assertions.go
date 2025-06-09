package main

import "fmt"

// Println is a helper function to print the arguments
func Println(a ...interface{}) {
	fmt.Println(a...)
}

// TypeAssertionsDemo demonstrates the usage of type assertions in Go
func TypeAssertionsDemo() {
	// In Go, interface{} is a special type that
	// - can hold values of *any* type
	// - Doesn't define any methods
	// - Acts like a container that can store any value
	// This is similar to Object in Java.
	var i interface{} = "Hello" // creates a variable of type interface{} that holds the value "Hello"

	// Where will we need this contraption? This pattern is useful when you need to:
	// - Write functions that can accept parameter of any type
	// - Store heterogeneous data in a collection
	// - Work with data whose type is not known at compile time

	// This is a type assertion and not a function call. It attempts to access the underlying value of i and checks if
	// that value is of type string.
	s, ok := i.(string)
	if ok {
		Println(s)
	} else {
		Println("Type assertion to String failed")
	}

	f, ok := i.(float64) // Type assertion to float64
	if ok {
		Println(f)
	} else {
		Println("Type assertion to float failed")
	}

	// In this example, i is an interface variable holding a string value. The first type assertion attempts to extract
	// the string value, which succeeds. The second type assertion attempts to extract a float64 value, which fails,
	//and the ok variable is set to false. It's crucial to check the ok value to avoid panics.
}
