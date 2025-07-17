package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ports := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// Start 10 workers
	for i := 0; i < cap(ports); i++ {
		go pworker(ports, results, &wg)
	}

	// Start go routine to collect and print results
	done := make(chan bool)
	go func() {
		var openPorts []int
		for port := range results {
			openPorts = append(openPorts, port)
		}

		// Print results in order
		for _, port := range openPorts {
			fmt.Printf("%d is open.\n", port)
			done <- true
		}
	}()

	// Send ports to workers
	for port := 1; port <= 300; port++ {
		wg.Add(1)
		ports <- port
	}

	wg.Wait()
	close(ports)
	close(results)
	<-done // Wait for the result to be printed.
}

func pworker(ports chan int, results chan int, wg *sync.WaitGroup) {
	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)

		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			wg.Done()
			continue
		}

		// send the result to the results channel instead of printing directly
		results <- port
		_ = conn.Close()
		wg.Done()

	}
}
