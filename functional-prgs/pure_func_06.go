package main

import (
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		accumulate(x)
	}
	os.Exit(y)
}

var y int // y is a package-level variable that accumulates the sum of integers passed as command-line arguments.

// accumulate adds the given integer x to the package-level variable y.
func accumulate(x int) {
	y += x
}
