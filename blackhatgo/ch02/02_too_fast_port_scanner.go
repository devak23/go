package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			fmt.Printf("trying %s\n", address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				return
			}
			_ = conn.Close()
			fmt.Printf("%d is open", j)
		}(i)

		// Without the WaitGroup, the code we just ran launches a single goroutine per connection, and the main
		// goroutine doesn’t know to wait for the connection to take place. Therefore, the code completes and exits as
		// soon as the for loop finishes its iterations, which may be faster than the network exchange of packets between
		// your code and the target ports. We will not get accurate results for ports whose packets were still in-flight.
		// One is to use WaitGroup from the sync package, which is a thread-safe way to control concurrency.
		// Once we’ve created WaitGroup, we can call a few methods on the struct.
		// The first is Add(int), which increases an internal counter by the number provided.
		// Next, Done() decrements the counter by one. Finally, Wait() blocks the execution of the goroutine in which
		// it’s called, and will not allow further execution until the internal counter reaches zero. This is shown
		// in "too_fast_port_scanner_fixed.go"

	}
}
