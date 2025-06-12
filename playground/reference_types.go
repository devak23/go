package main

import (
	"reflect"
)

// ReferenceTypesDemo this code demonstrates an important concept about Go slices and how they behave when passed to
// functions. The main point this file is trying to make is: Slices in Go are reference types, but when you pass a slice
// to a function, you're passing a copy of the slice header (which contains pointer, length, and capacity), not the
// slice itself.
func ReferenceTypesDemo() {
	slice := make([]int, 1, 7)                   // creates the slice with length 1 and capacity 7
	Println("slice1 ->", len(slice), cap(slice)) // prints it
	changeSlice(slice)                           // changes the slice's length and capacity inside the function
	Println("slice2 ->", len(slice), cap(slice)) // Still prints length 1 and capacity 7
	Println(slice)
}

// changeSlice function changes the slice's length and capacity inside the function. Inside the function the slice
// appears to have length of 5 and capacity of 5. This function uses reflection to bypass Go's normal slice semantics.
// Normally you can't modify the length or capacity of a slice directly, but reflection allows unsafe modifications of
// these internal headers. However, these changes only affect the local copy of the slice passed to the function. The
// original slice remains unchanged. This demonstrates that slices are "reference types" in that they point to the
// underlying data but slice header itself is copied when passed to functions.
func changeSlice(slice []int) {
	reflect.ValueOf(&slice).Elem().SetLen(5)     // &slice gets the address of the local slice parameter
	reflect.ValueOf(&slice).Elem().SetCap(5)     // reflect.ValueOf() creates a reflection object from the address: &slice
	Println("slice3 ->", len(slice), cap(slice)) // .Elem() deferences the pointer to get the actual slice value and
	Println(slice)                               // .SetLen(5), .SetCap(5) foribly modify the slice's length and capacity
}

// OUTPUT:
// slice1 -> 1 7
// slice3 -> 5 5
// [0 0 0 0 0]
// slice2 -> 1 7
// [0]
