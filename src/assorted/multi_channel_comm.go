package assorted

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, (x + y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// FiboMain is the main function thats invoked from main
func FiboMain() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // read from the channel
		}
		quit <- 0 // send zero to the quit channel
	}()

	fibonacci(c, quit)
}
