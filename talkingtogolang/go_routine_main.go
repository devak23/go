package main

// While the programs we write may do one thing and one thing really well, they may need to achieve this one thing by
// running small "pieces of work" in the background. This is done by writing asynchronous code. One of the easiest way
// to do this is using Go routines.

// Go routines are lightweight threads, processing units, that run concurrently with the main program.
// They are created using the `go` keyword, followed by a function call.
// They are great for running I/O operations, such as waiting for user input, or waiting for a response from a
// remote service.
// They are very fast to create, and they are very light-weight, so we can create millions of them.
// They are not preemptive, they will run until they call a blocking operation, or until they return.
// They are not scheduled by the operating system, they are scheduled by the Go runtime.
// They are not threads, they are goroutines.
// They are not processes, they are not isolated from each other.
// They can communicate with each other using channels.

import (
	"fmt"
	"time"
)

func GoRoutineMain() {
	go printNumbers("go-thread")
	printNumbers("main-thread")
}

func printNumbers(threadName string) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(threadName, "says: ", i)
	}
}
