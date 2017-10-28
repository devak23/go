package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const salutation = "Hello"

func main() {
	sayHello := func() {
		defer wg.Done()
		fmt.Println(salutation)
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
