package main

// This program concurrently executes HTTP GET requests against various search engines and prints out the amount of time
// each GET took.
import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	URL      string
	Duration time.Duration
}

// getter executes the HTTP GET request against url, prints out the duration of the call (or an error message if it failed),
// and calls wg.Done() when it finishes. It's meant to be run in a goroutine.
func getter(url string, ch chan<- result) {
	defer close(ch)     // Closes the channel when done.
	start := time.Now() //Records the start time.

	if _, err := http.Get(url); err != nil { //Tries to fetch the URL using http.Get.
		ch <- result{URL: url, Duration: time.Duration(-1)}
		return
	}

	ch <- result{URL: url, Duration: time.Since(start)} // sends the result with the URL and the duration it took to fetch it.
}

func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com", "http://www.duckduckgo.com"}
	durationChannel := make(chan result)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		_ = "breakpoint"
		ch := make(chan result) // create a new unbuffered channel for each URL
		go getter(url, ch)      // start the goroutine to fetch the URL and send it to the channel

		// this starts another goroutine that will read from the channel and forward the result to the durationChannel
		go func() {
			defer wg.Done()
			t := <-ch            // Waits to receive a result from the getter goroutine via ch
			durationChannel <- t // sends the result to the durationChannel for collection
		}()
	}

	// Cleanup goroutine
	go func() {
		wg.Wait()              // Wait for all workers to finish
		close(durationChannel) // Signal "no more data coming"
	}()

	// Collector can now safely range over the channel
	for t := range durationChannel {
		_ = "breakpoint"
		fmt.Println(t) // Process each duration
	}
	// Without the cleanup goroutine, the collector wouldn't know when to stop waiting for more data. Closing the
	// channel serves as a "end of stream" signal, allowing range loops to terminate gracefully.
}
