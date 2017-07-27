package waytogo

import (
	"fmt"
	"os"
	"bufio"
)

// ReadWithBufIoMain is the main function
func ReadWithBufIoMain() {
	// we declare the reference of the buffer
	var inputReader *bufio.Reader
	var input string
	var err error

	// we initialize the buffer by invoking the NewReader() and passing in
	// instance of Standard input
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a line: ")

	// ReadString()  will read a line of string till the delimiter is
	// encountered. The delimiter is also included in the line read from
	// the input. If the delimiter is not found an error is returned
	input, err = inputReader.ReadString('\n')

	if err  == nil {
		fmt.Printf("\nInput was: %s\n", input)
	}
}

