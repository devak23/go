package main

import . "fmt"

type Message struct {
	X string  // visible
	y *string // not visible
}

func (v *Message) Store(x, y string) {
	v.X = x
	v.y = &y
}

func (v *Message) Print() {
	if v.X != "" {
		Println(v.X, *v.y)
	} else {
		Println(v.X)
	}
}

func main() {
	m := &Message{}
	m.Print()
	m.Store("Hello", "World!")
	m.Print()
}
