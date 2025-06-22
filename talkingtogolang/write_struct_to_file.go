package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Salary struct {
	Basic float64
}

type MyEmployee struct {
	FirstName, LastName, Email string
	Age                        float64
	MonthlySalary              []Salary
}

func WriteStructIntoFileMain() {
	data := MyEmployee{
		FirstName: "Steven",
		LastName:  "King",
		Email:     "steven.king@example.com",
		Age:       50,
		MonthlySalary: []Salary{
			{Basic: 1000},
			{Basic: 2000},
			{Basic: 3000},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile("employee.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
