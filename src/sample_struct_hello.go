package main

import (
	"fmt"
	"container/list"
)

type info struct {
	result string
}

func infoRepo(aNumber int) (string, error) {
	return "This is an intro package", nil
}

func main() {
	var introMessage string = "Hello From GoToolChain"
	fmt.Println("Go Reports: %+v\n",introMessage)

	var info = info{}

	mssg, err := infoRepo(2)
	if err == nil {
		fmt.Println(mssg)
	}

	// structs define a flexible way of defining composite data types
	sp := &info

	sp.result = "set a struct pointer value"
	fmt.Println("Go reports: %+v\n", sp.result)

	xs := []float64 {98, 93, 77, 82, 83.5}
	total := 0.0
	for _, v := range xs  {
		total += v
	}
	fmt.Println("total = ", total)
	info.result = "This is a test string"
	fmt.Println("Another way of setting and reading values from a struct: ", info.result)

	l := list.New()
	l.PushBack(4)
	l.PushFront(1)
	l.PushBack(5)
	l.PushFront(6)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}