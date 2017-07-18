package assorted

import "fmt"

// -------------------Person delcaration------------------------
type NewPerson struct {
	name  string
	age   float32
	phone string
}

func (p *NewPerson) SayHi() {
	fmt.Printf("Hi! my name is %s\n", p.name)
}

// -------------------Employee declaration------------------------
type NewEmployee struct {
	NewPerson
	company string
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi! My name is %s and I'm an employee at %s\n", e.name, e.company)
}

// ------------------- Student declaration ------------------------
type NewStudent struct {
	NewPerson
	school string
}

func (s *Student) SayHi() {
	fmt.Printf("Hi! my name is %s and i study at %s\n", s.name, s.school)
}

// -----------------The main function --------------------------
func OOMthdOverloadMain() {
	mark := NewStudent{NewPerson{"Mark", 25, "333-444-3434"}, "MIT"}
	saul := NewEmployee{NewPerson{"Saul", 45, "769-342-8990"}, "Google"}

	mark.SayHi()
	saul.SayHi()
}
