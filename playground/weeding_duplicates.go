package main

import (
	"fmt"
	"github.com/samber/lo"
	"math/rand"
)

type Developer struct {
	Name string
	Age  int
}

func Unique(developers []Developer) []Developer {
	seen := make(map[Developer]bool)
	var uniques []Developer
	for _, developer := range developers {
		if !seen[developer] {
			seen[developer] = true
			uniques = append(uniques, developer)
		}
	}

	return uniques
}

func WeedingDuplicates() {
	var developers = []Developer{
		{Name: "Elliot", Age: rand.Intn(50)},
		{Name: "Allan", Age: 32},
		{Name: "Jennifer", Age: rand.Intn(60)},
		{Name: "Paul", Age: 30},
		{Name: "Graham", Age: rand.Intn(40)},
		{Name: "Paul", Age: 30},
		{Name: "Allan", Age: 32},
	}

	//uniques := Unique(developers)
	uniques := lo.Uniq(developers)

	fmt.Println(uniques)
}
