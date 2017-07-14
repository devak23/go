package gobyexample

import "fmt"

func ForMain() {

	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// an infinite for loop till you break out
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
