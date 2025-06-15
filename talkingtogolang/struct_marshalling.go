package main

import (
	"encoding/json"
	"fmt"
)

type Hello struct {
	Message string `json:"hellooo"`
}

func ReadingMarshalledStruct() {
	h := Hello{Message: "Hello World"}
	b, _ := json.Marshal(h)
	fmt.Println(string(b))
}
