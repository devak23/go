package main

import (
	. "fmt"
	"github.com/devak23/go/functional-prgs/utils"
)

type Person[T utils.Scalar] struct {
	Name string
	Age  T
}

func (p *Person[T]) SetName(name string) *Person[T] {
	return &Person[T]{
		Name: name,
		Age:  p.Age,
	}
}

func (p *Person[T]) SetAge(age T) *Person[T] {
	return &Person[T]{
		Name: p.Name,
		Age:  age,
	}
}

func (p *Person[T]) String() string {
	return "Name: " + p.Name + ", Age: " + Sprint(p.Age)
}

func main() {
	person := Person[float64]{
		Name: "Abhay",
		Age:  46.7,
	}
	anotherPerson := person.SetName("Guru").SetAge(41.5)

	Println("Original:", person.String())
	Println("Clone:", anotherPerson.String())
}
