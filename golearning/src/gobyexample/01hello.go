package gobyexample

import "fmt"

func HelloMain() {
	fmt.Println("======= HelloMain =========")
	fmt.Println("Hello World")

	var title string
	var copies int

	title = "For the love of Go"
	copies = 1000

	fmt.Println(title, copies)
}
