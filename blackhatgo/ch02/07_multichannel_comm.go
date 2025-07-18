package main

// To solve the problem in prior implementations, we need to use a separate thread to pass the result of the port scan
// back to our main thread to order the ports before printing. Another benefit of this modification is that it removes
// the dependency of a WaitGroup entirely, as we’ll have another method of tracking completion. For example, if we scan
// 1024 ports, we’re sending on the worker channel 1024 times, and we’ll need to send the result of that work back to
// the main thread 1024 times. Because the number of work units sent and the number of results received are the same,
// our program can know when to close the channels and subsequently shut down the workers.

import (
	"fmt"
	"net"
	"sort"
)

// The portWorker(ports, results chan int) function has been modified to accept two channels. The remaining logic is
// mostly the same, except that if the port is closed, we’ll send a zero , and if it’s open, we’ll send the port
func portWorker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		fmt.Println("trying ", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		_ = conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int) // create a separate channel to communicate the results from the worker to the main thread
	var openPorts []int       // to store the results so we can sort them later

	for i := 0; i <= cap(ports); i++ {
		go portWorker(ports, results)
	}

	// Start feeding the ports into the ports channel. We need to send to the workers in a separate goroutine because
	// the result-gathering loop needs to start before more than 100 items of work can continue.
	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	// The result-gathering loop receives on the results channel 1024 times as the worker thread either sends in the
	// port number or a zero. If the port doesn’t equal 0, it’s appended to the slice
	for i := 0; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open \n", port)
	}
}

// The higher the count of the workers, the faster the program should execute. But if we add too many workers, our
// results could become unreliable. When we’re writing tools for others to use, we’ll want to use a healthy default
// value that caters to reliability over speed. However, we should also allow users to provide the number of workers
// as an option.
