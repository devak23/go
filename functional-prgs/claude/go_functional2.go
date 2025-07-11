package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

type PersonStats struct {
	AverageAge float64
	Count      int
	Cities     []string
}

func IsAgeMoreThan25Years(p Person) bool {
	return p.Age >= 25
}

func peopleGreaterThan25Years(people []Person) []Person {
	return Filter(people, IsAgeMoreThan25Years)
}

func processIntoStats(people []Person) []Person {
	// group by city and calculate stats
	cityGroup := make(map[string][]Person)
	for _, p := range people {
		cityGroup[p.City] = append(cityGroup[p.City], p)
	}

	// log processing
	fmt.Printf("Processing %d people in %d cities\n", len(people), len(cityGroup))
	return people
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func complexFunctionalExample() {
	fmt.Println("\n=== People, City & Stats ===")

	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "San Francisco"},
		{"Charlie", 35, "New York"},
		{"Diana", 28, "Chicago"},
		{"Eve", 32, "San Francisco"},
	}

	stats := Pipeline(people, peopleGreaterThan25Years, processIntoStats)

	// Calculate final statistics
	totalAge := Reduce(stats, 0, func(acc int, p Person) int { return acc + p.Age })
	avgAge := float64(totalAge) / float64(len(stats))
	cities := Map(stats, func(p Person) string { return p.City })
	uniqueCities := removeDuplicates(cities)

	finalStats := PersonStats{
		AverageAge: avgAge,
		Count:      len(stats),
		Cities:     uniqueCities,
	}

	fmt.Println("Final stats: ", finalStats)
}
