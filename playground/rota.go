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
		endDate := advanceDays(startDate, 4)
		lead := pickALead(leads, i)
		devs := joinNames(rotateLeft(&developers)[:3])

		fmt.Printf("%v - %v, %v, %v\n", formatShortDate(startDate), formatShortDate(endDate), devs, lead)
		startDate = advanceDays(endDate, 3)
	}
}

func formatShortDate(t time.Time) string {
	return t.Format("01/02")
}

func rotateLeft(developers *[]string) []string {
	first := (*developers)[0]
	*developers = append((*developers)[1:], first)
	return *developers
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
