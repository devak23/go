package goinpractise

import (
	"fmt"
	"time"
)

// Count counts and prints the output with a slight delay
func Count() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

// ConcurrentPrintingMain is the main function to be invoked from the main package
func ConcurrentPrintingMain() {
	// invoke the ro-routine in a separate thread
	go Count()
	// sleep for 2 milliseconds
	time.Sleep(time.Millisecond * 2)
	// then print Hello Go Routine
	fmt.Println("Hello Go Routine")
	// and then again sleep for 5 milliseconds
	time.Sleep(time.Millisecond * 5)
}
