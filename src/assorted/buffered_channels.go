package assorted

import "fmt"

// BufferedChannelsMain makes use of buffered channels.
func BufferedChannelsMain() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // uncommenting this line causes a run time exception
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
