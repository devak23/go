package concurrency

import (
	"fmt"
	"golearning/src/goreceipes/concurrency/syncutils"
	"time"
)

// this is to make the main program wait for other go routines to finish

// the main program
func PassingTheBuckMain() {

	// create a channel that will be shared with 2 go routines
	channel := make(chan int)
	// set the waitGroup to 2 so that the program waits for 2 go routines before terminating
	syncutils.Wg.Add(2)

	fmt.Println("Starting Goroutines")
	// start the two go routines
	go passTheBuck("routine-1", channel)
	go passTheBuck("routine-2", channel)

	fmt.Println("Communication between channel begins")

	// initiate the data flow
	channel <- 1 // <-- this is a blocking call. Main go routine will be blocked here
	// wait for the goroutines to finish
	syncutils.Wg.Wait()
	fmt.Println("\nTerminating program")
}

func passTheBuck(chName string, ch chan int) {
	// Schedule the WaitGroup's Done method
	defer syncutils.Wg.Done()
	for {
		// receive message from the channel
		value, ok := <-ch
		// check if the channel was closed before you process anything
		if !ok {
			fmt.Printf("channel was closed for %s\n", chName)
			return
		}

		// Print the value received from the channel
		fmt.Printf("Count %d received from go routine %s\n", value, chName)
		// simulate a delay
		time.Sleep(time.Duration(1 * time.Second))

		// if the value is 10, then close the channel
		if value == 10 {
			fmt.Printf("Channel closed from %s\n", chName)
			close(ch)
			return
		}
		// else increment the counter
		value++

		// send the value back to another go routine
		ch <- value // <-- This is a blocking call
	}
}

// You can see that the buck gets passed between routine1 and routine2. The reason that happens is because
// the main goroutine does not read the channel and thus has no role to play in incrementing the counter.
// Therefore when routine 1 writes into the channel, there is only routine 2 which reads the value.
