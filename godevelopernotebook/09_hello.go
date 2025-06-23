package main

import . "fmt"

type HelloWorld struct {
}

func (h HelloWorld) String() string {
	return "Hello World!"
}

type Message struct {
	HelloWorld
}

func main() {
	m := Message{}
	Println("via struct: ", m.HelloWorld.String())
	Println("invoking String on struct: ", m.String())
	Println("printing struct itself: ", m)
}
