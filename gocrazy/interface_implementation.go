package main

import (
	"fmt"
	"math/rand"
)

type AllSafeEmployee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Python"
}

func (e Engineer) Age() int {
	return rand.Intn(40)
}

func TypesAgainDemo() {
	elliot := Engineer{Name: "Elliot Alderson"}
	var programmers []AllSafeEmployee
	programmers = append(programmers, elliot)
	fmt.Println(programmers)
}

// What this program demonstrates is that AllSafeEmployee is an interface that is implemented by the struct Engineer
// via 2 methods on line# 17 and 21. Remove these 2 functions, and you will have trouble appending elliot to the
// AllSafeEmployee[] slice on line# 28
