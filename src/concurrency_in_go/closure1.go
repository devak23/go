package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// this is how we define a closure. Much similar to Javascript.
func main() {
	salutation := "Hello"
	
	sayHello := func() {
		defer wg.Done()
		fmt.Println(salutation)
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
