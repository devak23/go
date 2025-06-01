package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// here, the expectation is that the program will print out hello, greetings and good day!
// one by one, but it doesn't happen so. The reason is the main goroutine (having the control of
// the CPU runs the assignment much faster than any of the goroutines executing thereafter).
// thus the result you see is - good day! good day! good day!
func main() {
	for _, salutation := range []string{"hello", "greetings", "good day!"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
