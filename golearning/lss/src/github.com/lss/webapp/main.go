package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	fmt.Println("A simple File server")
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "public"+r.URL.Path)
// 	})

// 	fmt.Println("file server running on port 8000...")
// 	http.ListenAndServe(":8000", nil)
// }

// even more elegant is the following:

func main() {
	fmt.Println("File server listening on port 8000...")
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
