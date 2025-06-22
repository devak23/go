package main

import (
	"fmt"
	"log"
	"time"
)

// A go routine frequently connects with a Go channel. Channels allow different go routines to communicate easily and efficiently.

// The following program shows how the data is sent to a channel and how it is read from it

func GoChannelMain() {
	log.Println("------------ GoChannelMain: Demo of channel creation -------------------")
	log.Println("Creating a channel")
	c := make(chan int)
	log.Println("Writing numbers into the channel in a go routine")
	go PrintNumbersToChannel(c)

	log.Println("Printing numbers in main")
	for num := range c {
		fmt.Println(num)
	}
}

func PrintNumbersToChannel(c chan int) {
	for i := 0; i < 10; i++ {
		log.Printf("Writing %d to the channel %v\n", i, c)
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
	log.Println("Closing channel")
	close(c) // Closing the channel is a good practice as it signals that no more data will be sent and helps prevent deadlocks.
}

// Note that we printed into the channel in a goroutine, but we read it in the main in the for loop. So how was this possible?
// This is because for (range in c) is a blocking operation. It will keep on reading from the channel until the channel is closed.
// If we didn't have this for loop, the main function would have ended and the program would have exited before the go routine could
// write to the channel. Now that begs the question that channel is being closed in PrintNumbersToChannel(), so why would it be closed
// there and not in main? - Well the reason is that the fundamental principle in Go is that goroutine that writes to a channel is
// responsible for closing it. PrintNumbersToChannel() is the sender and GoChannelMain() is the receiver. So it's the sender's job
// to close the channel.
// Why closing in main would be problematic: ?
// If you moved close(c) to the main function, you'd face these issues:
// 1. Race condition: Main might close the channel while PrintNumbersToChannel() is still trying to write to it, causing a panic
// 2. Timing complexity: Main would need to somehow know when the sender is done writing
// 3. Violates Go conventions: The sender should control the channel lifecycle
// So how would one read from a channel in a go routine? For that, please refer: read_from_channel.go
