package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("A simple File server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			fmt.Printf("Uh Oh... an error! %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		} else {
			// copy the contents of the file directly onto the Response Writer
			// ReponseWriter implements the Writer interface, this is ok.
			io.Copy(w, f)
		}
		defer f.Close()
	})

	fmt.Println("file server running on port 8000...")
	http.ListenAndServe(":8000", nil)
}
