package concurrency

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

// a wait group is used make the program wait till goroutines are done with executions
var wg sync.WaitGroup

func WaitGroupMain() {
	// add a count of two to each goroutine
	wg.Add(2)

	fmt.Print("Starting Goroutines")

	go addTable()
	go multiplyTable()

	// wait for goroutines to finish
	fmt.Println("Waiting for goroutines to finish execution...")
	wg.Wait()
	fmt.Println("\nTerminating the MAIN goroutine")
}

func addTable() {
	// Schedule a call to WaitGroup's Done method to tell goroutine is completed
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// generate a random duration duration
		duration := rand.Int63n(1000)
		// sleep for those many milliseconds
		time.Sleep(time.Duration(duration) * time.Millisecond)
		// now do the addition table
		fmt.Println("Addition Table for: ", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d + %d = %d\t", i, j, i+j)
		}
		fmt.Println("\n")
	}
}

func multiplyTable() {
	// Schedule call to WaitGroup's Done method to tell that goroutine has completed its execution
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// generate a random delay duration
		duration := rand.Int63n(1000)
		// sleep for that duration
		time.Sleep(time.Duration(duration) * time.Millisecond)
		// now do the multiplication table
		fmt.Println("Multiplication Table for: ", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, j*i)
		}
		fmt.Println("\n")
	}
}
