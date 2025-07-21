package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// The Response type contains an exported Body parameter, which is of type io.ReadCloser. An io.ReadCloser is an
	// interface that acts as an io.Reader as well as an io.Closer, or an interface that requires the implementation of
	// a Close() function to close the reader and perform any cleanup. The details are somewhat inconsequential; just
	// know that after reading the data from an io.ReadCloser, we’ll need to call the Close() function on the response
	// body. Using defer to close the response body is a common practice; this will ensure that the body is closed
	// before we return it.

	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	// Print HTTP status
	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)

	// Read and display response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(body))
}

// Courtesy Claude -
// OLD (deprecated)          			NEW (Go 1.16+)
// ioutil.ReadAll(r)          			→ io.ReadAll(r)
// ioutil.ReadFile(name)      			→ os.ReadFile(name)
// ioutil.WriteFile(name, data, perm) 	→ os.WriteFile(name, data, perm)
// ioutil.ReadDir(name)       			→ os.ReadDir(name)
// ioutil.TempFile(dir, pattern) 		→ os.CreateTemp(dir, pattern)
// ioutil.TempDir(dir, pattern)  		→ os.MkdirTemp(dir, pattern)
// ioutil.Discard            			→ io.Discard
// ioutil.NopCloser(r)        			→ io.NopCloser(r)
