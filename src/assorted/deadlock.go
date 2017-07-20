package assorted

import "fmt"

func DeadlockMain() {
	channel := make(chan int)
	// This will create a deadlock.
	channel <- 10 // we are sending data to the channel
	fmt.Println(<- channel) // and we are also reading from it
	// Since the send operation is done in the main goroutine, it blocks the execution of
	// the execution of the rest of the flow. The send operation always waits for somebody
	// to receive at the other end. Since that's not happening, you shouldn't see the following
	// println
	fmt.Println("Done!")
}
