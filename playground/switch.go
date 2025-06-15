package main

import "time"

func SwitchDemo() {
	timeNow := time.Now()

	day := timeNow.Day()
	switch {
	case day >= 1 && day <= 5:
		Println("We are in the beginning of the month")
	case day >= 6 && day <= 20:
		Println("We are in mid-month")
	default:
		Println("We are in the fag end of the month")
	}

	switch timeNow.Weekday() {
	case time.Saturday, time.Sunday:
		Println("It's the weekend!", timeNow.Weekday())
	default:
		Println("It's a weekday.", timeNow.Weekday())
	}
}
