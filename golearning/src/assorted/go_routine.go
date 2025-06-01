package assorted

import (
	"fmt"
	"runtime"
)

func sayIt(s string) {
	for i := 0; i < 5; i++ {
		// this yields the processor so that other goroutines
		// can run. It does not suspend the current goroutine
		// so the execution resumes automatically
		runtime.Gosched()
		fmt.Println(s)
	}
}

func GoRoutineMain() {
	go sayIt("World") // create a new goroutine
	sayIt("Hello")    // main goroutine
}
