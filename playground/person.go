package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Friend interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("Hello", p.Name)
}

func Greet(f Friend) {
	f.SayHello()
}

// Demonstrates two ways of creating an object.
func main() {
	var guy = Person{Name: "Dave", Age: 32}
	guy.SayHello()

	var anotherGuy = new(Person)
	anotherGuy.Name = "Melnek"
	anotherGuy.Age = 60
	anotherGuy.SayHello()

	Greet(&guy)       // notice how I have to pass in the & (address) or pointer to the Greet function
	Greet(anotherGuy) // Here I don't have to because new(Person) returns a pointer to Person not a value of type Person

	var dog = new(Dog)
	Greet(dog) // this is called duck-typing as Greet takes in anyone which implements the Friend interface
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof! woof!")
}
