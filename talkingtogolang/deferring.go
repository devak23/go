package main

import "fmt"

func DeferingDemo() {
	printDeferring()
	deferringFunctions()
}

func printDeferring() {
	defer fmt.Println("This is deferred")
	fmt.Println("This is not deferred")
}

func deferringFunctions() {
	defer func() {
		fmt.Println("This will be printed in the end")
	}()
	defer fmt.Println("This will be printed second")
	fmt.Println("This will be printed first!")
}

// The central idea conveyed by this code is to demonstrate Go's defer statement and its LIFO (Last In, First Out) execution order.
// Defer keyword defers the execution of a function to the end of the current function
// The defer statement serves several important purposes:
// 1. Guaranteed execution: Deferred functions execute when the surrounding function returns, regardless of how it returns (normal return, panic, etc.)
// 2. LIFO order: Multiple deferred statements execute in Last-In-First-Out order, like a stack
// 3. Resource cleanup: Perfect for closing files, releasing locks, or cleaning up resources
// 4. Code organization: Keeps cleanup code near the resource allocation code

// The deferred statements execute in reverse order after the regular fmt.Println completes, demonstrating the stack-like
// behavior that makes defer reliable for cleanup operations.
