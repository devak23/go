package main

import "flag"
import "net/http"

func main() {
	//flag takes 3 arguments the name of the flag, its default value and the description.

	port := flag.String("port", "3000", "Port to serve HTTP on")
	path := flag.String("path", ".", "The directory to serve")

	flag.Parse()

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*path)))
}

// usage: go run fs.go -port=5151 -path="/tmp"
