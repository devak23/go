package main

import "fmt"

func main() {
	greeting := greet(EN)
	fmt.Println(greeting)
}

// We define a type language based on String, but Go identifies language and string as different types.
type language string

// You can attach methods on language but not on String for instance:
func (l language) IsValid() bool {
	return l == EN || l == FR
}

// you can define constants such as
const (
	EN language = "en"
	FR language = "fr"
)

// ... which is a better type check than String that is passed on freely ("en" or "fr" may contain typos). Besides, it's
// a self-documenting code as the greet function below takes in a specific type rather than String
func greet(l language) string {
	switch l {
	case EN:
		return "Hello, World"
	case FR:
		return "Bonjour, le Monde"
	default:
		return ""
	}
}
