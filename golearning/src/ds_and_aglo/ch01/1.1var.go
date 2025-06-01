package main

import "fmt"

func main() {
	var v1 int
	var v2 int
	v1 = 100

	fmt.Println("Value stored in variable v1::", v1)
	fmt.Println("Value stored in variable v2::", v2)

	var truth bool = true
	fmt.Println("truth ::", truth)

	if !truth {
		fmt.Println("What's the point of living?")
	}

}
