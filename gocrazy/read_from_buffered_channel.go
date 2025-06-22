package main

import (
	"log"
	"sync"
)

func ReadFromBufferedChannelMain() {
	log.Println("--- ReadFromBufferedChannelMain: Demo of reading from a buffered channel ---")
	c := make(chan int, 15)
	go PrintNumbersToChannel(c)
	// No need for time.Sleep() here as the buffer allows writer to continue
	var wg sync.WaitGroup
	wg.Add(1)
	go ReadNumbersFromChannel(c, &wg)
	wg.Wait()
}

// Unbuffered channel (make(chan int)):
//	1. Has no internal storage
//	2. Sender blocks until a receiver is ready
//	3. Receiver blocks until a sender sends data
//	4. Synchronous communication
// Buffered channel (make(chan int, capacity)):
//	1. Has internal storage for capacity values
//	2. Sender only blocks when buffer is full
//	3. Receiver only blocks when buffer is empty
//	4. Asynchronous communication (up to buffer capacity)

// With a buffered channel, the writer goroutine can send up to 15 values before blocking, making the communication more
// efficient in scenarios where the producer and consumer operate at different speeds.
// This is because the buffer acts as a temporary storage, allowing the writer to continue without waiting for the reader
// to process the data, as long as the buffer has capacity.
// If the buffer is full, the writer will block until there is space available.
// The reader can process values from the buffer at its own pace, without blocking the writer, as long as the buffer is
// not empty. If the buffer is empty, the reader will block until the writer adds new values.
