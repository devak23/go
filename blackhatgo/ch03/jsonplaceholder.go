package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func (a Address) String() string {
	return fmt.Sprintf("\n%s,\n%s, %s\n%s", a.Suite, a.Street, a.City, a.Zipcode)
}

type Person struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s\nUsername: %s\nAddress: %s\nEmail: %s",
		p.Name, p.Username, p.Address, p.Email)
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var person Person
	if err := json.NewDecoder(resp.Body).Decode(&person); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("%s", &person)

}
