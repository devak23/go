package gobyexample

import "fmt"

func valuesMain() {

	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("(1+1) again with brackets  = ", (1 + 1))
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	//fmt.Println(true & false) & operator is not defined on booleans
}
