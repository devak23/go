package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}

	const numOfRuns = 5
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < numOfRuns; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
