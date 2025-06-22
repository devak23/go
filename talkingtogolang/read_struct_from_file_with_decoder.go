package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type AnotherEmployee struct {
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	EmployeeId int       `json:"employee_id"`
	ManagerId  int       `json:"manager_id"`
	HireDate   time.Time `json:"hire_date"`
}

func ReadStructFromFileWithDecoderMain() {
	jsonFile, err := os.Open("employees.json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatalf("Error reading employees json: ", err)
	}

	decoder := json.NewDecoder(jsonFile)

	// Read the opening bracket
	_, err = decoder.Token()
	if err != nil {
		log.Fatalf("Error reading opening bracket: ", err)
	}

	// Process each employee
	var employees []AnotherEmployee
	for decoder.More() {
		var employee AnotherEmployee
		err := decoder.Decode(&employee)
		if err != nil {
			log.Fatalf("Error decoding employee: ", err)
		}
		employees = append(employees, employee)
	}

	fmt.Println(employees)
}
