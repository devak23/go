package main

import . "fmt"

// HelloWorld defines a struct (class) with no members in it
type HelloWorld struct {
}

// String() - method is being defined on the struct HelloWorld
func (h HelloWorld) String() string {
	return "Hello World!"
}

// Message - is another struct being defined which contains an anonymous field of type HelloWorld. This is called type
// embedding using an anonymous field. Go’s design has upset quite a few people with an inheritance-based view of object
// orientation because it lacks inheritance, however thanks to type embedding we’re able to compose types which act as
// proxies to the methods provided by anonymous fields
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
