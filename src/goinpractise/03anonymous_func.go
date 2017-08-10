package goinpractise

import (
	"fmt"
	"runtime"
)

// AnonymousMain is the function that gets called from main
func AnonymousMain() {
	fmt.Println("Outside a goroutine...")

	// run an anonymous function inside a go routine
	go func() {
		fmt.Println("Inside a goroutine")
	}()

	fmt.Println("Again outside the goroutine");

	// yield the control to the scheduler so that goroutine gets a chance of
	// doing something. If this step isn't present, then the main program terminates
	// without giving a chance to the goroutine to execute its task.
	runtime.Gosched()
}
