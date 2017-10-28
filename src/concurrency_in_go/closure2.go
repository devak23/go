package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

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
