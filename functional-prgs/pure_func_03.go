package main

import (
	"github.com/devak23/go/functional-prgs/utils"
	"os"
)

func main() {
	// command-line arguments are passed as 5 and 6. Hence, exit code is 11
	os.Exit(utils.AddScalar(utils.Arg(0), utils.Arg(1)))
}
