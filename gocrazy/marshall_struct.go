package main

import (
	"encoding/json"
	"fmt"
)

type Hello struct {
	// You can specify metadata on fields of structs using what is called a tag line. This tag line is used for
	// different things. One common use is to format the output of the fields in JSON. That tag line can also be used to
	// format data for persistence to database, for example: to write a tag line by adding a specific directive after
	// the field’s type, using backquotes, as shown below.
	Message string `json:"hellooo"`

	// what it says is:
	// Message string - defines a struct field named "Message" of type string
	// `json:"hellooo"` - is a struct tag that tells the JSON encoder/decoder to use "hellooo" as the JSON key name instead of "Message"
	// So when you marshal this struct to JSON, instead of getting:
	// {"Message": "Hello World"}
	// you will get:
	// {"hellooo": "Hello World"}

	// The struct tag allows you to customize how Go struct fields are represented in JSON (or other formats), enabling
	// you to use different naming conventions or map to existing JSON schemas that don't match your Go field names
}

func MarshallStructMain() {
	h := Hello{Message: "Hello World"}
	b, _ := json.Marshal(h)
	fmt.Println(string(b))
}
