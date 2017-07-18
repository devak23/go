package assorted

import "fmt"

type Person struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Person // anonymous field
	school string
}

type Employee struct {
	Person  // anonymous field
	company string
}

// define a method in Person
func (p *Person) SayHi() {
	fmt.Printf("Hi, I'm %s. You can call me on %s.\n", p.name, p.phone)
}

// OOMthdInteritanceMain gets invoked from main.go What this example shows
// is once you define a method for parent struct, it becomes available
// to whichever child struct that the parent struct composes of. In this case
// the Person struct's method is inherited by the Student and Employee structs
func OOMthdInteritanceMain() {
	saul := Student{Person{"Saul", 25, "620-923-8989"}, "MIT"}
	steve := Employee{Person{"Steve", 40, "510-924-3434"}, "Google"}

	saul.SayHi()
	steve.SayHi()

}
