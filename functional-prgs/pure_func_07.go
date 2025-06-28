package main

import (
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a.Add(x)
	}

	os.Exit(int(a))
}

var a Accumulator // a is a package-level variable of type Accumulator that represents sum of integers

type Accumulator int

func (a *Accumulator) Add(x int) {
	*a += Accumulator(x)
}
