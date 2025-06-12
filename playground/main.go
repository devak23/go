package main

import "fmt"

// Println is a helper function to print the arguments
func Println(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	UsingMapAsSet()
	UsingRecommendationSystem()
	TypeAssertionsDemo()
	ReferenceTypesDemo()
	InterfaceTypesDemo()
}
