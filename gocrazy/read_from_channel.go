package main

import (
	"log"
	"sync"
	"time"
)

func ReadFromChannelMain() {
	log.Println("--- Reading from a channel ---")
	c := make(chan int) // create an unbuffered integer channel
	go PrintNumbersToChannel(c)
	time.Sleep(2 * time.Second)
	// Pausing for 2 seconds before the reader starts. During this time the writer goroutine is actively writing the numbers
	// to the channel. The numbers accumulate in the channel (though it's unbuffered, so writer blocks after the first
	// write)

	// Main wakes up and creates a WaitGroup to coordinate with the reader coroutine
	var wg sync.WaitGroup
	// it adds 1 to the counter indicating we are waiting for 1 coroutine to complete
	wg.Add(1)

	// Main spawns another goroutine to read from the channel and passes the pointer to the WaitGroup
	go ReadNumbersFromChannel(c, &wg)
	// Wait for the reader to finish

	wg.Wait()
}

func ReadNumbersFromChannel(c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // ensures that the WaitGroup is decremented when the function exits
	// reads the number from the channel and prints it out
	for n := range c {
		log.Printf("Number received: %d\n", n)
	}
}

// Key Concepts Demonstrated:
// 1. Goroutine Communication: Writer and reader communicate through the channel
// 2. Synchronization: WaitGroup ensures main waits for reader completion
// 3. Channel Lifecycle: Writer closes channel, reader detects closure and exits
// 4. Blocking Operations: Channel operations block until data is available or channel is closed
// The program effectively demonstrates reading from a channel in a separate goroutine while maintaining proper
// synchronization with the main function.
