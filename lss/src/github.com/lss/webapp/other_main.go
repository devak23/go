package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func setContentType(w http.ResponseWriter, urlPath string) {
	var contentType string
	switch {
	case strings.HasSuffix(urlPath, "css"):
		contentType = "text/css"
	case strings.HasSuffix(urlPath, "html"):
		contentType = "text/html"
	case strings.HasSuffix(urlPath, "png"):
		contentType = "image/png"
	default:
		contentType = "text/plain"
	}
	w.Header().Add("Content-Type", contentType)
}

func main1() {
	fmt.Println("A simple File server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			fmt.Printf("Uh Oh... an error! %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		} else {
			defer f.Close()
			// The handler also need to tell the browser what kind of file it is
			setContentType(w, r.URL.Path)

			// copy the contents of the file directly onto the Response Writer
			// ReponseWriter implements the Writer interface, this is ok.
			io.Copy(w, f)

		}
	})

	fmt.Println("file server running on port 8000...")
	http.ListenAndServe(":8000", nil)
}
