package main

import (
	"fmt"
	"time"
)

// A go routine frequently connects with a Go channel. Channels allow different go routines to communicate easily and efficiently.

// The following program shows how the data is sent to a channel and how it is read from it

func GoChannelMain() {
	c := make(chan int)
	go printNumbersToChannel(c)

	for num := range c {
		fmt.Println(num)
	}
}

func printNumbersToChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}
