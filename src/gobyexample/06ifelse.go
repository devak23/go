package main

import "fmt"


// There is no ternary if statement in GO

func main() {
	// parenthesis are not required 
	// for if conditions but braces are a MUST
	if 7 % 2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	// if without an else
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// you can assign and compare on the same lines
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "consists of single digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
