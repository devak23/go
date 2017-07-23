package concurrency

import (
	"goreceipes/concurrency/syncutils"
	"fmt"
	"math"
)

// create a struct to hold the number and it's processed value
type NumberObject struct {
	number int
	value  int
}

// Main
func MultiChannelCommMain() {
	// notify the main of 4 threads in play
	syncutils.Wg.Add(4)

	// create channels for each function
	squareCh := make(chan NumberObject)
	fibCh := make(chan NumberObject)
	doubleCh := make(chan NumberObject)

	// launch threads to calculate values
	go calculateSquares(squareCh)
	go calculateFibonacci(fibCh)
	go calculateDouble(doubleCh)

	// launch the printer thread
	go printChannels(squareCh, fibCh, doubleCh)

	// wait for threads to complete
	syncutils.Wg.Wait()

	fmt.Println("Terminating program.")
}

// print the output of each channel
func printChannels(sqCh <-chan NumberObject, fibCh <-chan NumberObject, doubleCh <-chan NumberObject) {
	defer syncutils.Wg.Done()
	exhausedMap := make(map[string]int) // maintains a count of how many channels are exhausted
	exhausedMap["sqCh"] = 0
	exhausedMap["fibCh"] = 0
	exhausedMap["doubleCh"] = 0
	for i:= 0; i<30; i++ {
		//if exhausedMap["sqCh"] == 1 && exhausedMap["fibCh"] == 1 && exhausedMap["doubleCh"] == 1{
		//	fmt.Println("All channels are exhausted")
		//	return
		//}
		select {
		case sqObj, ok := <-sqCh:
			if ok {
				fmt.Printf("Square of %d = \t%d\n", sqObj.number, sqObj.value)
			} else {
				exhausedMap["sqCh"] = 1
			}
		case fibObj, ok := <-fibCh:
			if ok {
				fmt.Printf("Fibonacci of %d = %d\n", fibObj.number, fibObj.value)
			} else {
				exhausedMap["fibCh"] = 1
			}
		case doubleObj, ok := <-doubleCh:
			if ok {
				fmt.Printf("Double of %d = \t%d\n", doubleObj.number, doubleObj.value)
			} else {
				exhausedMap["doubleCh"] = 1
			}
		}
	}
}

// calculates double
func calculateDouble(doubleCh chan<- NumberObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		doubleCh <- NumberObject{number: i, value: i * 2}
	}
}

// calculate fibonacci
func calculateFibonacci(fibCh chan<- NumberObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		num := float64(i)
		Phi := (1 + math.Sqrt(num)) / 2
		phi := (1 - math.Sqrt(num)) / 2
		result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
		fibCh <- NumberObject{number: int(num), value: int(result)}
	}
}

// calculates squares
func calculateSquares(sqCh chan<- NumberObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		sqCh <- NumberObject{number: i, value: i * i}
	}
}
