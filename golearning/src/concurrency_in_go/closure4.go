package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// now we pass the string got from iterating the slice into the closure
func main() {
	for _, salutation := range []string{"hello", "greetings", "good day!"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
