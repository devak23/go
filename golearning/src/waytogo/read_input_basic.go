package waytogo

import "fmt"

// ReadInputBasicMain is the main function
func ReadInputBasicMain() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)

	fmt.Printf("Please enter your fullname: ")

	// Scanln scans the text input from the keyboard storing the space
	// separated values into the respective variables
	fmt.Scanln(&firstName, &lastName)

	// fmt.Scanf("%s %s", &firstName, &lastName) // Scanf also does the same but uses the format parameters as the first argument
	fmt.Printf("\nHi %s %s!\n", firstName, lastName)

	// Sscanf and the related functions do the same but they dont read from the standard input
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From string we read: ", f, i, s)
}
