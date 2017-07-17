package assorted

import (
	"fmt"
	"os"
)

// ConvertStringMain is the main function that gets called from main.go
func ConvertStringMain() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <name>")
		os.Exit(1)
	}

	var name = os.Args[1]
	fmt.Printf("You entered: %s\n", name)
	// name[0] = 'C' // uncommenting this line causes compilation error
	// you cannot change the string in Go. Strings are immutable. The only
	// way to change a string is to create a new string as follows
	byteArray := []byte(name)
	byteArray[0] = 'C'
	changedName := string(byteArray)
	fmt.Printf("Changed name: %s\n", changedName)
}
