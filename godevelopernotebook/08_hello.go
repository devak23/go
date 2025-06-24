package main

import . "fmt"

// Message defines a struct as an area of allocated memory which is subdivided into slots for holding named values,
// where each named value has its own type. Here we’ve defined a struct Message which contains two values: X and y
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
		Println("Printing when not empty: ", v.X, *v.y)
	} else {
		Println("Printing when empty: ", v.X)
	}
}

// Both methods above have a pointer receiver (*Message). In Go, when you call a method with a pointer receiver on a
// value, Go will automatically take the address of the value for you. Therefore in the main program below, it doesn't
// matter if you initialize m with &Message or not. However, in the case below where a Person struct is being initialized
// with or without a reference(pointer) type, since the methods do not accept a pointer receiver, a different variable
// is created when the method is invoked and is destroyed when the method exits. Therefore, even if you pass the pointer
// to the person struct into these methods, the output still gives you zero!

func main() {
	m := &Message{}
	m.Print()
	m.Store("Hello", "World!")
	m.Print()

	p := &Person{}
	p.Print()
	p.Store("Abhay", 46)
	p.Print()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {
	Println("Printing Person attribs: ", p.Name, p.Age)
}

func (p Person) Store(name string, age int) {
	p.Name, p.Age = name, age
}
