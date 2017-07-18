package assorted

import "fmt"

func addOne(a *int) *int {
	*a = *a + 1
	return a
}

// PassByRefMain is the main function invoked from main.go
func PassByRefMain() {
	x := 12
	fmt.Println("x = ", x)

	x1 := addOne(&x)
	fmt.Println("x = ", x)
	fmt.Println("x1 = ", *x1)
}
