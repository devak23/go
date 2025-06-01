package goinpractise

import (
	"fmt"
	"net/http"
)

// hello is the handler of the request
func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World! %s", request.URL.Path[1:])
}

// HelloServerMain gets invoked from the main.go. After running this program
// access it from the browser: http://localhost:4000/
func HelloServerMain() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
	fmt.Println("Server listening on port 4000...")
}
