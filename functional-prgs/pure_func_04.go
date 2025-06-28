package main

import (
	"github.com/devak23/go/functional-prgs/utils"
	"os"
	"strconv"
)

func main() {
	// strconv.Atoi converts a string to an integer.
	// os.Args[0] is the program name, so we start from os.Args[1] for the first argument.
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	os.Exit(utils.AddScalar(x, y))
}
