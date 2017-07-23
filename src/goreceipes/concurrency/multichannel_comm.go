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

type QuitObject struct {
	channelName string
	quitValue int
}

// Main
func MultiChannelCommMain() {
	// notify the main of 4 threads in play
	syncutils.Wg.Add(4)

	// create channels for each function
	squareCh := make(chan NumberObject)
	fibCh := make(chan NumberObject)
	dblCh := make(chan NumberObject)
	quitCh := make(chan QuitObject, 3)

	// launch threads to calculate values
	go calculateSquares(squareCh, quitCh)
	go calculateFibonacci(fibCh, quitCh)
	go calculateDouble(dblCh, quitCh)

	// launch the printer thread
	go printChannels(squareCh, fibCh, dblCh, quitCh)

	// wait for threads to complete
	syncutils.Wg.Wait()

	fmt.Println("Terminating program.")
}

// print the output of each channel
func printChannels(sqCh <-chan NumberObject, fibCh <-chan NumberObject, dblCh <-chan NumberObject, quitCh <- chan QuitObject) {
	defer syncutils.Wg.Done()
	exhausedMap := make(map[string]int) // maintains a count of how many channels are exhausted
	exhausedMap["sqCh"] = 0
	exhausedMap["fibCh"] = 0
	exhausedMap["dblCh"] = 0
	for {
		select {
		case obj, _ := <-sqCh:
			fmt.Printf("Square of %d = \t%d\n", obj.number, obj.value)
		case obj, _ := <-fibCh:
			fmt.Printf("Fibonacci of %d = %d\n", obj.number, obj.value)
		case obj, _ := <-dblCh:
			fmt.Printf("Double of %d = \t%d\n", obj.number, obj.value)
		case val := <- quitCh:
			exhausedMap[val.channelName] = val.quitValue
			if exhausedMap["sqCh"] == 1 && exhausedMap["fibCh"] == 1 && exhausedMap["dblCh"] == 1{
				fmt.Println("All channels are exhausted")
				return
			}
		}
	}
}

// calculates double
func calculateDouble(dblCh chan<- NumberObject, quitCh chan <- QuitObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		dblCh <- NumberObject{number: i, value: i * 2}
	}
	// send the quit signal
	quitCh <- QuitObject{"dblCh", 1}
}

// calculate fibonacci
func calculateFibonacci(fibCh chan<- NumberObject, quitCh chan <- QuitObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		num := float64(i)
		Phi := (1 + math.Sqrt(num)) / 2
		phi := (1 - math.Sqrt(num)) / 2
		result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
		fibCh <- NumberObject{number: int(num), value: int(result)}
	}
	// send the quit signal
	quitCh <- QuitObject{"fibCh", 1}
}

// calculates squares
func calculateSquares(sqCh chan<- NumberObject, quitCh chan <- QuitObject) {
	defer syncutils.Wg.Done()

	for i := 0; i < 10; i++ {
		sqCh <- NumberObject{number: i, value: i * i}
	}
	// send the quit signal
	quitCh <- QuitObject{"sqCh", 1}
}
