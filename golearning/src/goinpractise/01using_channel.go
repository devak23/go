package goinpractise

import (
	"fmt"
	"time"
)

// PrintCount accepts an int type channel
func PrintCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c // waits for the value to come in
		fmt.Print(num, " ")
	}
}

// UsingChannelMain is the main function that will get called from main.go
func UsingChannelMain() {
	// create a channel
	c := make(chan int)
	// data array
	a := []int{8, 6, 7, 5, 3, 1, 2, 0, -1}

	// start a go routine
	go PrintCount(c)

	for _, value := range a {
		c <- value // passes int values into the channel
	}
	// The main thread pauses before ending
	time.Sleep(time.Millisecond * 1)
	fmt.Println("end of main")
}
