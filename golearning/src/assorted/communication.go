package assorted

import (
	"fmt"
	"time"
)

func CommunicationMain() {
	buffChannel := make(chan int, 3) // synchronous operation
	unbuffChannel := make(chan int)  // asynchronous operation

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		unbuffChannel <- 1
		close(unbuffChannel)
		// note that the channel was closed in the goroutine itself
		// ... because it's a synchronous operation
	}()

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		buffChannel <- 3
		time.Sleep(time.Duration(1) * time.Second)
		buffChannel <- 4
		time.Sleep(time.Duration(1) * time.Second)
		buffChannel <- 5
		time.Sleep(time.Duration(1) * time.Second)
	}()

	// Print the values from the unbuffered channel
	fmt.Println("waiting for values...")
	fmt.Println(<-unbuffChannel)
	// try to read from the closed channel
	val, ok := <-unbuffChannel
	if ok {
		fmt.Println("Still more data! ->", val) // this wont execute
	}

	// Print the values from the buffered channel
	fmt.Println(<-buffChannel)
	fmt.Println(<-buffChannel)
	fmt.Println(<-buffChannel)
	close(buffChannel)
	// note the channel was closed in the main goroutine because it's asynchronous operation

	// try to read from the closed channel
	val, ok = <-buffChannel
	if ok {
		fmt.Println("Channel is closed.. .but there is more data! ->", val) // this wont execute
	}
}

// also note, there is no need for waitgroup here because we are relying on the behavior
// of the channel and the fact that channel is being shared by the 2 goroutines with the
// main goroutine. IF there was no sharing of data with the main goroutine, we would have
// had to use the waitgroup for the main goroutine to pause till the others finished their
// job
