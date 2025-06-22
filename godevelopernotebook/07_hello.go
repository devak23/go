package main

import . "fmt"

const Hello = "Hello"

var myWorld = "world"

func main() {
	myWorld := myWorld + "!"
	Println(Hello, myWorld)

	var w string
	w = myWorld + "!"
	Println(Hello, w)

	Println(Hello, world())

	Println(message("Gopher!"))

	Println(greet())

	welcome("Welcome to Go!", "Let's learn together!", "Hello World!", "I am liking Go!")
}

func world() string {
	return "World!"
}

func message(name string) string {
	return Sprintf("Hello %v\n", name)
}

func greet() (string, string) {
	return "Hello", "World!"
}

func welcome(v ...interface{}) {
	Println(v...)
}
