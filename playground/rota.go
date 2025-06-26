package main

import (
	"fmt"
	"strings"
	"time"
)

func RotaMain() {
	developers := []string{"Bhaskar", "Varkha", "Nimisha", "Dhiren", "Abrar", "Nitin"}
	leads := []string{"Abhay", "Saif"}

	startDate := time.Date(2025, time.June, 23, 0, 0, 0, 0, time.FixedZone("IST", 5*60*60+30*60))

	for i := 0; i < 10; i++ {
		// Calculate the end date (Friday) as 4 days after the start date
		endDate := advanceDays(startDate, 4)
		// Select the lead for the week based on the index
		lead := pickALead(leads, i)
		// Rotate the developers to get the next set of developers for the week
		rotatedDevs := rotateLeft(developers)

		fmt.Printf("%v - %v, %v, %v\n", formatShortDate(startDate), formatShortDate(endDate), joinNames(rotatedDevs[:3]), lead)
		// Advance the start date by 3 days for getting the next Monday
		startDate = advanceDays(endDate, 3)
		// assign the rotated developers back to the original slice for the next iteration
		developers = rotatedDevs
	}
}

// formatShortDate formats the time.Time object to a short date string in the format "MM/DD".
// Go's time formatting uses the unique reference date (Jan 2, 2006, 15:04:05 MST) to assign meaning to numbers in format
// strings. "01/02" is always interpreted as "month/day" because 01 is the month and 02 is the day in the reference date.
func formatShortDate(t time.Time) string {
	return t.Format("01/02")
}

func rotateLeft(developers []string) []string {
	first := developers[0]
	developers = append(developers[1:], first)
	return developers
}

func joinNames(slice []string) string {
	return strings.Join(slice, ", ")
}

func pickALead(leads []string, i int) string {
	return leads[(i/2)%len(leads)]
}

func advanceDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// Output:
// 06/23 - 06/27, Varkha, Nimisha, Dhiren, Abhay
// 06/30 - 07/04, Nimisha, Dhiren, Abrar, Abhay
// 07/07 - 07/11, Dhiren, Abrar, Nitin, Saif
// 07/14 - 07/18, Abrar, Nitin, Bhaskar, Saif
// 07/21 - 07/25, Nitin, Bhaskar, Varkha, Abhay
// 07/28 - 08/01, Bhaskar, Varkha, Nimisha, Abhay
// 08/04 - 08/08, Varkha, Nimisha, Dhiren, Saif
// 08/11 - 08/15, Nimisha, Dhiren, Abrar, Saif
// 08/18 - 08/22, Dhiren, Abrar, Nitin, Abhay
// 08/25 - 08/29, Abrar, Nitin, Bhaskar, Abhay
