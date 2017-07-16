package gplch01

import (
	"fmt"
	"os"
	"strings"
)

/*
JoinArgsMain is a simple routine that joins the arguments using strings.Join function
*/
func JoinArgsMain() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	for index, value := range os.Args {
		fmt.Printf("%d has the parameter %s\n", index, value)
	}
}
