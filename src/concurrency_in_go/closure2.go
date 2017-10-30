package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// here we shadow the salutation variable declared in the main block and see that 
// go routines indeed work in the same memory space as of the main go routine because
// the variable value is changed to World
func main() {
	salutation := "Hello"
	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "World"
	}()
	wg.Wait()
	fmt.Println(salutation)
}
