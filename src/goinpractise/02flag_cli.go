package goinpractise

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to")
var spanish bool
var russian bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
	flag.BoolVar(&russian, "russian", false, "Use Russian language")
	flag.BoolVar(&russian, "r", false, "Use Russian language")
}

// FlagCliMain is the main function called from main.go
// Run the application using the command go run src/main/main.go -name "Buttercup" --spanish
func FlagCliMain() {
	flag.Parse() // parse the flags placing the values in variables
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else if russian == true {
		fmt.Printf("Здравствуйте %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
