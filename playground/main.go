package main

import (
	"fmt"
	"runtime"
)

// Println is a helper function to print the arguments
func Println(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	version := runtime.Version()
	fmt.Printf("Go version: %s\n", version)

	WeedingDuplicates()
	UsingMapAsSet()
	UsingRecommendationSystem()
	TypeAssertionsDemo()
	ReferenceTypesDemo()
	InterfaceTypesDemo()
	SwitchDemo()
}
