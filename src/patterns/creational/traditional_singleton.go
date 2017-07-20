package creational

import (
	"sync"
	"sync/atomic"
)

// define the ElvisSingleton interface
type ElvisSingleton interface {
	AddOne() int
}

// define a Struct whose singleton will be created
type elvis struct {
	// it contains the count variable which keeps track
	// of the number of instances
	count int
}

// create a global reference which will point to a singleton
var instance *elvis

// an atomic uint32 flag that tells if the instance was initialized or not
var initialized uint32

// create a synchronization object so that the creational code can lock
var monitor sync.Mutex

// GetElvis will be invoked using the pointer of ElvisSingleton
func GetElvis() *elvis {
	// check if the instance is already initialized. This is being done by checking the atomic
	// variable and returning the instance if it is initialized
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// acquire the lock (synchronization)
	monitor.Lock()
	defer monitor.Unlock()

	// following the double checked locking paradigm
	if initialized == 0 {
		//instance = &elvis{1} OR
		instance = new(elvis)  // ---> actually create the instance

		// set the flag to "initialized"
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// provide the implementation of the interface into the instance
func (e *elvis) AddOne() int {
	e.count++
	return e.count
}
