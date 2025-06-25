package main

import "fmt"

type HelloWorld bool

func (h HelloWorld) String() (r string) {
	if h {
		r = "Hello World!"
	}
	return
}

type Message struct {
	HelloWorld
}

func main() {
	m := &Message{HelloWorld: true}
	fmt.Println("true by default: ", m.String())
	m.HelloWorld = false
	fmt.Println("set to false: ", m.String())
	m.HelloWorld = true
	fmt.Println("set to true: ", m.String())
}
