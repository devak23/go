package episode00

// This program implements a thread-safe in-memory hash table (key-value store) using built-in map with proper
// concurrency control.

import (
	"sync"
)

type inMemoryHashTable struct {
	data map[string][]byte
	lock sync.RWMutex // A reader-writer mutex for thread safety
}

// NewInMemoryDb - What's interesting in this implementation is that the compiler has no problem returning an instance
// of inMemoryHashTable as an object of type "HashTable" (interface defined in hashtable.go). This is because, in Go,
// a type automatically satisfies an interface if it implements all the methods that the interface declares. The compiler
// sees that inMemoryHashTable has methods with exact signatures matching the HashTable interface, and therefore it can
// be returned. Thus key points are:
// 1. No explicit "implements" keyword needed (unlike Java/C#)
// 2. Duck typing: "If it walks like a duck and quacks like a duck, it's a duck"
// 3. Compile-time checking: The compiler verifies method signatures match exactly
// 4. Pointer receivers: Since methods have pointer receivers (ht *inMemoryHashTable)

func NewInMemoryDb() HashTable {
	return &inMemoryHashTable{data: make(map[string][]byte)}

	// Another point to note is that we initialized the data but we didn't initialize the mutex. The mutex doesn't need
	// explicit initialization because of Go's zero value concept. A zero-value sync.RWMutex is immediately usable. One
	// could explicitly initialize the mutex, but it's unnecessary
	//
	// return &inMemoryHashTable {
	//    data: make(map[string][]byte),
	//    lock: sync.RWMutex{},  // Same as zero value and therefore redundant.
	// }
	// Go's designers made synchronization primitives like sync.Mutex and sync.RWMutex work immediately upon creation,
	// so you don't need to explicitly initialize them. This is different from maps, slices, and channels which need
	// make() to be usable.
}

func (ht *inMemoryHashTable) Get(key string) ([]byte, error) {
	ht.lock.RLock()         // for read access - multiple goroutines can read simultaneously
	defer ht.lock.RUnlock() // ensures the lock is released when the function exits

	val, ok := ht.data[key] // checks for the existence of the key and returns value if present else returns error
	if !ok {
		return val, ErrorNotFound
	}
	return val, nil
}

func (ht *inMemoryHashTable) Set(key string, value []byte) error {
	ht.lock.Lock()         // for write access - exclusive access, blocks all other operations
	defer ht.lock.Unlock() //ensures the lock is released

	ht.data[key] = value // assigns the value to the map and returns nil as a success operation
	return nil
}

// Key design decisions:
// 1. Thread safety: The RWMutex allows multiple concurrent readers but exclusive writers, which is optimal for
// read-heavy workloads
// 2. Byte slices: Values are stored as []byte rather than strings, making it more flexible for storing any kind of data
// 3. Error handling: Uses Go's idiomatic error return pattern
// 4. Defer statements: Ensures locks are always released, even if the function panics
