package main

import (
	"fmt"
	"golearning/src/programmingingo/stacker"
)

func main() {
	var haystack stacker.Stack

	// pushing the elements on the stack
	haystack.Push("hay")
	haystack.Push(-16.5)
	haystack.Push(20)
	haystack.Push(1 == 1)
	haystack.Push([]string{"pin", "clip", "neeldle"})
	haystack.Push("stack")

	// poping the elements from stack
	for {
		item, err := haystack.Pop()
		if err != nil {
			fmt.Printf("%v\n", err)
			break
		}
		fmt.Println(item)
	}
}
