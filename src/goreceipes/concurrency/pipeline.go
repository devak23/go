package concurrency

import (
	"math/rand"
	"math"
	"fmt"
)

// This program demonstrates a 3 stage pipe line with 3 goroutines and 2 channels connecting them. First
// thread will simply generate random numbers and send it to the next goroutine. The second goroutine
// will compute Fibonacci for that number and pass it to the next goroutine. The third goroutine
// simply prints the numbers as it receives
func PipelineMain() {
	// tell the main goroutine that there are 3 goroutines to be waited for
	waitGroup.Add(3)

	// create channels to be shared between the goroutines.
	chRandomNumbers := make(chan int) // shared between 1st and 2nd
	chFibsForNumbers := make(chan fibvalue) // shared between 2nd and 3rd
	// launch goroutines now...

	// this thread will simply generate random numbers
	go generateRandomNumber(chRandomNumbers)

	// this thread will generate fibonacci numbers for those random numbers generated
	go generateFibonacci(chFibsForNumbers, chRandomNumbers)

	// this thread will simply print the fibonacci numbers of those random numbers
	go printFibonacci(chFibsForNumbers)

	// wait for all goroutines
	waitGroup.Wait()
}

// the function that randomly generates random numbers and sends it out on the channel
// to goroutine# 2
func generateRandomNumber(out chan int) {
	defer waitGroup.Done()
	var random int
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random
	}
	// close the channel as the other goroutines are blocked.
	close(out)
}

// we need a structure that will hold the index and it's fibonacci value
type fibvalue struct {
	index, fiboValue int
}

// generateFibonacci will read the values from the "in" channel, compute the Fibonacci
// value using the Binet's formula and write to the out channel to goroutine #3
func generateFibonacci(out chan fibvalue, in chan int) {
	// let the main goroutine know that this goroutine is done
	defer waitGroup.Done()
	var input float64
	for v := range in {
		input = float64(v)
		// Calculate Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)

		// shove the value and the index into the out channel to be fed to goroutine# 3
		out <- fibvalue{
			index:     v,
			fiboValue: int(result),
		}
	}
	// close the channel as it's a blocking call.
	close(out)
}

// printFibonacci will read the channel and print each fiboancci numbers
func printFibonacci(in chan fibvalue) {
	// let the main goroutine know this goroutine is done executing.
	defer waitGroup.Done()

	// print the values of the number and it's fiboanacci value
	for v := range in {
		fmt.Printf("Fibonacci fiboValue of %d is %d\n", v.index, v.fiboValue)
	}
}
