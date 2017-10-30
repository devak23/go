package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// This is an example where the Producer is less active than the numerous consumers that
// the code creates.
func main() {
	// define a producer which simply counts 5 seconds actively getting and releasing the lock
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i >= 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	// define a consumer that which does nothing but simply acquires and reliquishes the lock
	consumer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	// now define a test function that will use both the producer and consumer
	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		//fmt.Println("(test) ==> count", count)
		wg.Add(count + 1)

		beginTestTime := time.Now()
		// start the producer routine
		go producer(&wg, mutex)
		// start 'n' number of consumers, all of them sharing the same lock as the producer.
		for i := count; i > 0; i-- {
			go consumer(&wg, rwMutex)
		}

		// wait for routines to get over and return the time taken
		wg.Wait()
		return time.Since(beginTestTime)
	}

	// code to print the results on the screen and also to execute tests
	// in which
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var rwMutex sync.RWMutex
	fmt.Fprintf(tw, "Readers\trwMutex\tMutex\n")
	for i := 0; i <= 20; i++ {
		count := int(math.Pow(2, float64(i)))
		//fmt.Println("count = ", count)
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &rwMutex, rwMutex.RLocker()),
			test(count, &rwMutex, &rwMutex))
	}
}
