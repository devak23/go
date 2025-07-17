package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// To avoid inconsistencies, we’ll use a pool of goroutines to manage the concurrent work being performed. Using a for
// loop, we’ll create a certain number of worker goroutines as a resource pool. Then in the main, we will use channel
// to provide work

func main() {
	// First, you create a channel by using make(). A second parameter, an int value of 100, is provided to make() here.
	// This allows the channel to be buffered, which means you can send it an item without waiting for a receiver to read
	// the item. Buffered channels are ideal for maintaining and tracking work for multiple producers and consumers.
	// We’ve capped the channel at 100, meaning it can hold 100 items before the sender will block. This is a slight
	// performance increase, as it will allow all the workers to start immediately.
	ports := make(chan int, 200)
	var wg sync.WaitGroup

	// Next, we use a for loop to start the desired number of workers in this case, 100.
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, &wg)
	}

	// Iterating over the ports sequentially we send a port on the ports channel to the worker.
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
}

// worker takes two arguments: a channel of type int and a pointer to a WaitGroup. The channel will be used to receive
// work, and the WaitGroup will be used to track when a single work item has been completed.
func worker(ports chan int, wg *sync.WaitGroup) {

	// we use range to continuously receive from the ports channel, looping until the channel is closed.
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second) // we added a DialTimeout with a 2 second wait if we are not able to make connections

		if err != nil { // if there is no connection, then the worker is not killed,
			wg.Done()
			continue
		}
		fmt.Printf("%d is open.\n", p)
		_ = conn.Close()
		wg.Done()
	}
}
