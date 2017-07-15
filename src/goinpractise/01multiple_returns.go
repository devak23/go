package goinpractise

// import the FormattedIO package
import (
	"fmt"
)

// MultipleReturns returns multiple return values. This is a public function
func MultipleReturns() (string, string) {
	return "Abhay", "Soham"
}

// MultipleReturnsWithParams accepts two string arguments and returns them
// The example only illustrates that multiple return values can be returned
// from the function
func MultipleReturnsWithParams(firstname string, lastname string) (string, string) {
	return firstname, lastname
}

// MutlipleReturnsMain is the main function called MultipleReturnsMain so as to be
// invoked from main.go
func MutlipleReturnsMain() {
	n1, n2 := MultipleReturns()
	fmt.Println(n1, n2)

	// anything with an underscore is ignored
	n3, _ := MultipleReturnsWithParams("Manik", "Kulkarni")
	fmt.Println(n3)
}
