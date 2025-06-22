package main

import (
	"log"
	"sync"
)

func ReadFromBufferedChannelMain() {
	log.Println("--- ReadFromBufferedChannelMain: Demo of reading from a buffered channel ---")
	c := make(chan int, 15)
	go PrintNumbersToChannel(c)

	var wg sync.WaitGroup
	wg.Add(1)
	go ReadNumbersFromChannel(c, &wg)
	wg.Wait()
}
