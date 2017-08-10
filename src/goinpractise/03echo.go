package goinpractise

import (
	"fmt"
	"io"
	"os"
	"time"
)

func EchoMain() {
	// start a thread that copies anything from stdin to stdout
	go echo(os.Stdin, os.Stdout)

	// wait for 30 seconds
	time.Sleep(time.Second * 30)

	// print out a message saying "exiting"
	fmt.Println("Terminating the program")

	// exit the program
	os.Exit(0)
}

// echo copies everything from the reader object to the writer
func echo(in io.Reader, out io.Writer) {
	fmt.Println("Enter a line of text and see it echo back...")
	// Copy function requires the desination first and
	// then the source
	io.Copy(out, in)
}
