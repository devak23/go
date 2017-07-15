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
func MultipleReturnsWithParams(name1 string, name2 string) (string, string) {
	return name1, name2
}

// MutlipleReturnsMain is the main function called MultipleReturnsMain so as to be
// invoked from main.go
func MutlipleReturnsMain() {
	n1, n2 := MultipleReturns()
	fmt.Println(n1, n2)

	// anything with an underscore is ignored
	n3, _ := MultipleReturnsWithParams("Manik", "Soham")
	fmt.Println(n3)
}
