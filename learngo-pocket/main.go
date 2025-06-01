package main

import "fmt"

func main() {
	greeting := greet(EN)
	fmt.Println(greeting)

	phGreeting := greetWithPhrasebook(UR)
	fmt.Println(phGreeting)
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
	HE language = "he"
	UR language = "ur"
	DE language = "de"
	VI language = "vi"
	EL language = "el"
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

// However, as the languages grow, the switch case isn't very readable.To trim down our function without losing any of
// its functionality, we can introduce a map that maps languages to greetings.
var phrasebook = map[language]string{
	EN: "Hello, World",
	FR: "Bonjour, le Monde",
	HE: "שלום עולם",
	EL: "γεια σου κόσμε",
	UR: "ہیلو دنیا",
	VI: "xin chào thế giới",
	DE: "Hallo Welt",
}

func greetWithPhrasebook(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language %q", l)
	}
	return greeting
}
