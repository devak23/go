package main

import "fmt"

func strlen(s string, c chan int) {
	c <- len(s) // here the data is flowing into the channel
}

func main() {
	c := make(chan int) // define a channel which will take integers
	go strlen("Salutations", c)
	go strlen("World", c) // notice the channel used in both cases is the same.
	x, y := <-c, <-c      // this operator shows if the data is flowing to or from the channel. In this case, the data is flowing from the channel into the variable
	fmt.Printf("x = %d, y = %d, x+y = %d\n", x, y, x+y)
}
