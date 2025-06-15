package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Hello   string
	ignored string
}

func (m Message) String() string {
	return m.Hello + m.ignored
}

func StructDemo() {
	message := Message{Hello: "Hello World!", ignored: "ignored"}
	fmt.Printf("Message: %s\n", message.Hello)
	fmt.Printf("Mess 1: %s\n", message)
	fmt.Printf("Mess 2: %v\n", message)
	fmt.Printf("Mess 3: %+v\n", message)

	AsString, _ := json.Marshal(message)
	fmt.Printf("Mess 4: %s\n", AsString)
}
