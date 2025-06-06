package main

import (
	"encoding/json"
	"flag" // import the flag package
	"fmt"
	"os"
)

func main() {
	// The flag package offers two very similar functions to read a string from the command line. The first requires a
	// pointer to a variable that it will fill up:

	var lang string
	flag.StringVar(&lang, "lang", "en", "language to greet")

	// The second creates the pointer and returns it:
	// lang := flag.String("lang", "en", "language to greet")

	flag.Parse()
	// this scans the input parameters and fills every variable we’ve identified as a receiver. You can run this program
	// from the command line with the -lang flag:
	// go run main.go -lang=en

	greeting := greet(language(lang))
	fmt.Println("Step 1: With if else condition:", greeting)

	phGreeting := greetWithPhrasebook(language(lang))
	fmt.Println("Step 2: With a hardcoded phrase book: ", phGreeting)

	// In the 3rd stage we read a json file which contains a list of languages and their greetings.
	file, err := os.Open("dictionary.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// The file will be automatically closed when the function returns.
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&phrasebook)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
	phGreeting = greetWithCustomPhraseBook(language(lang), phrasebook)
	fmt.Println("Step 3: By reading a JSON file: ", phGreeting)
}

// We define a type language based on String, but Go identifies language and string as different types.
type language string

// You can attach methods on language but not on String for instance:
func (l language) IsValid() bool {
	return l == EN || l == FR
}

// The (l language) represents the receiver of the IsValid method. In other words, (l language) is the handle of the type
// language on which you can invoke the IsValid method.

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

func greetWithCustomPhraseBook(l language, phrasebook map[language]string) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language %q", l)
	}
	return greeting
}
