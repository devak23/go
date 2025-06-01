package assorted

import "fmt"

// MapDemoMain gets invoked from main.go
func MapDemoMain() {
	// initializing a map
	rating := map[string]float32{"C": 2, "Java": 5.0, "Go": 4.5, "Python": 4.5, "C++": 2}

	// map has two return values
	goRating, ok := rating["Go"]
	if ok {
		fmt.Printf("Go has a rating of %f\n", goRating)
	} else {
		fmt.Println("Go does not have any rating")
	}

	delete(rating, "Java")

	// print the remaining map
	for key, value := range rating {
		fmt.Printf("%s = %f\n", key, value)
	}
}
