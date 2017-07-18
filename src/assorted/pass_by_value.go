package assorted

import "fmt"

func add1(a int) int {
	a = a + 1
	return a
}

// PassByValueMain is the main function invoked from main.go
func PassByValueMain() {
	x := 3
	fmt.Println("x = ", x) // this will print 3

	x1 := add1(x)
	fmt.Println("x = ", x)   // this will print 3
	fmt.Println("x1 = ", x1) // this will print 4
}
