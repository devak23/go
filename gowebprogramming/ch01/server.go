package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	bytes, err := fmt.Fprintf(writer, "Hello %s!", request.URL.Path[1:])
	fmt.Printf("Request for path: %s\n", request.URL.Path)
	fmt.Println("# of bytes written: ", bytes)
	if err != nil {
		_ = fmt.Errorf("error writing to the response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
