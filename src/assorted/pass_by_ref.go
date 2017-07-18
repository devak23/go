package assorted

import "fmt"

func addOne(a *int) *int { // <---- note the differences here
	*a = *a + 1
	return a
}

// another way to define the function using references
func addTwo(a *int) int { // <---- note the differences here
	*a = *a + 2
	return *a
}

// PassByRefMain is the main function invoked from main.go
func PassByRefMain() {
	x := 12
	fmt.Println("x = ", x)

	x1 := addOne(&x)
	fmt.Println("x = ", x)
	fmt.Println("x1 = ", *x1) // <---- note the differences here

	y := 25
	fmt.Println(" y =", y)

	y1 := addTwo(&y)
	fmt.Println("y1 = ", y1)
	fmt.Println("y = ", y) // <---- note the differences here
}
