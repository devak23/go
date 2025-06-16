package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Employee struct {
	EmployeeId        int       `json:"employee_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	HireDate          time.Time `json:"hire_date"`
	JobId             string    `json:"job_id"`
	Salary            float64   `json:"salary"`
	CommissionPercent float64   `json:"commission_percent"`
	ManagerId         int       `json:"manager_id"`
	DepartmentId      int       `json:"department_id"`
}

func ReadStructFromFile() {
	jsonFile, _ := os.Open("./employees.json")
	defer jsonFile.Close()
	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	var employees []Employee
	_ = json.Unmarshal(jsonBytes, &employees)
	fmt.Println(employees)
}
