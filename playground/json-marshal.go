package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Customer struct {
	Name          string
	AccountNumber string
}

func main() {
	c := Customer{"John Doe", "4864089691879077"}
	newC := Customer{}
	jsonBytes, _ := json.Marshal(c)
	fmt.Println(string(jsonBytes))
	_ = json.Unmarshal(jsonBytes, &newC)
	fmt.Println(newC)

	f := Foo{"Joe Junior", "Hello Shabado"}
	newF := Foo{}
	xmlBytes, _ := xml.Marshal(f)
	fmt.Println(string(xmlBytes))
	_ = xml.Unmarshal(xmlBytes, &newF)
	fmt.Println(newF)
}

type Foo struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}
