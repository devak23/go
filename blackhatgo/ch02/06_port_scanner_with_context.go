package main

import (
	"context"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

func main() {
	// Configuration - easy to adjust
	const (
		numWorkers  = 20
		timeout     = 30 * time.Second
		connTimeout = 2 * time.Second // Longer connection timeout
		startPort   = 1
		endPort     = 1000
	)

	fmt.Printf("Scanning ports %d-%d with %d workers...\n", startPort, endPort, numWorkers)
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ports := make(chan int, numWorkers*2) // Buffer size based on workers
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go portScanner(ctx, ports, results, &wg, connTimeout)
	}

	// Result collector with sorting - READ IMMEDIATELY, don't wait for close
	var openPorts []int
	var mu sync.Mutex
	done := make(chan bool)

	go func() {
		for port := range results {
			fmt.Printf("Found open port: %d\n", port)
			mu.Lock()
			openPorts = append(openPorts, port)
			mu.Unlock()
		}
		done <- true
	}()

	// Send work
	wg.Add(endPort - startPort + 1) // Add all work upfront
	go func() {
		defer close(ports)
		for port := startPort; port <= endPort; port++ {
			select {
			case ports <- port:
				// wg.Add(1) moved outside - already added all work
			case <-ctx.Done():
				fmt.Println("\nOperation cancelled or timed out")
				return
			}
		}
	}()

	wg.Wait()
	close(results)
	<-done

	// Sort results for cleaner output
	sort.Ints(openPorts)

	fmt.Printf("\nFound %d open ports:\n", len(openPorts))
	for _, port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}

	fmt.Printf("\nScan completed in %v\n", time.Since(start))
}

func portScanner(ctx context.Context, ports chan int, results chan int, wg *sync.WaitGroup, connTimeout time.Duration) {
	for {
		select {
		case port, ok := <-ports:
			if !ok {
				return
			}

			address := fmt.Sprintf("127.0.0.1:%d", port)

			// More aggressive timeout for faster scanning
			dialer := &net.Dialer{Timeout: connTimeout}
			conn, err := dialer.DialContext(ctx, "tcp", address)
			if err != nil {
				wg.Done()
				continue
			}

			select {
			case results <- port:
			case <-ctx.Done():
				conn.Close()
				wg.Done()
				return
			}

			conn.Close()
			wg.Done()

		case <-ctx.Done():
			return
		}
	}
}
