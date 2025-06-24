package main

import . "fmt"

// HelloWorld defines a struct (class) with no members in it
type HelloWorld struct {
}

// String() - method is being defined on the struct HelloWorld
func (h HelloWorld) String() string {
	return "Hello World!"
}

// Message - is another struct being defined which contains an anonymous field of type HelloWorld
type Message struct {
	HelloWorld
}

// main invokes String() on the Message which calls String() on the embedded type HelloWorld. It can also be invoked directly
// on the Message type. Line# 24 and 25 are essentially the same.
func main() {
	m := Message{}
	Println("via struct: ", m.HelloWorld.String())
	Println("invoking String on struct: ", m.String())
	Println("printing struct itself: ", m)
}
