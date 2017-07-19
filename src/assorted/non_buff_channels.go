package assorted

import "fmt"

// loops over the array and puts the sum into the channel
func sum(a []int, channel chan int) {
	total := 0
	for _, value := range a {
		total += value
	}
	channel <- total // insert into the channel
}

// NonBufferedChannels is the main function that gets invoked from main.go
func NonBufferedChannels() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	channel := make(chan int)

	// start from the beginning to the mid
	go sum(array[:len(array)/2], channel)

	// start from the mid to the end
	go sum(array[len(array)/2:], channel)

	// now receive from the channel
	x, y := <-channel, <-channel

	fmt.Printf("X = %d\nY = %d\nX+Y = %d\n", x, y, (x + y))
}
