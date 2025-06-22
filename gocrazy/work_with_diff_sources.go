package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// To act on values coming from different channels, Go offers a "select block" where you can act depending on the value
// read from different async channels

// The `select` blocks untile one of its cases is receiving messages. In the following program, the main receives messages
// from two goroutines and then receives a message from time.After inline, then times out.

func WorkingWithDiffSourcesMain() {
	log.Println("----------------WorkingWithDiffSourcesMain: Combining Hello and World written by 2 goroutines  ----------------------")

	workingWithForloop()
	workingWithWaitGroup()
	workingWithTimeouts()
}

func workingWithWaitGroup() {
	log.Println("@@@ Working with WaitGroup @@@")

	ch := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	// Start readers before writers
	go func() {
		for msg := range ch {
			fmt.Print(msg)
		}
	}()

	// Start writers
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("Hello")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf(" WaitGroup!")
	}()

	wg.Wait()
	close(ch)
}

func workingWithForloop() {
	log.Println("- ---- --- Working with Forloop --- ----- -")
	ch := make(chan string)

	hello := func() {
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("Hello")
	}

	world := func() {
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf(" Forloop!")
	}

	go func() {
		defer close(ch)
		hello()
		world()
	}()

	for msg := range ch {
		fmt.Printf(msg)
	}
	fmt.Println("")
}

func workingWithTimeouts() {
	log.Println("~~~~ Working with Timeouts ~~~~")
	ch := make(chan string)

	helloThread := func() {
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("Hello")
	}
	go helloThread()

	worldThread := func() {
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf(" Timeout!")
	}
	go worldThread()

	for {
		select {
		case msg := <-ch:
			fmt.Printf("%s", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("\nWaited for 3 seconds")
			os.Exit(0)
		}
	}
}

// Why is the timeout needed in the select clause?
// Because the program will hang indefinitely if the timeout is not set. The select block is a blocking construct and
// go's way of multiplexing channel's operation allowing to handle multiple concurrent data sources. It will wait for
// any of the channels to receive a message. If none of the channels receive a message, the program will keep waiting.
// The timeout is a safety mechanism to prevent the program from hanging if one of the channels never receives a message.

// Alternate way of doing the same thing or removing the timeout is to use a for loop or using a wait group.
