package main

import (
	"github.com/devak23/go/functional-prgs/utils"
	"os"
	"strconv"
)

func main() {
	var sum int
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		sum = utils.Add(sum, x)
	}
	os.Exit(sum)
}
