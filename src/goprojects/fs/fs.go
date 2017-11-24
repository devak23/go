package main

import "os"
import "net/http"

func main() {
	dir, _ := os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}
