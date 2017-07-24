package assorted

import (
	"sync"
	"fmt"
	"net/http"
	"io/ioutil"
)

var waitGroup sync.WaitGroup

func FetchUrlsMain() {
	var urls = []string {
		"http://www.yahoo.com",
		"http://www.msn.com",
		"http://www.nytimes.com",
		"http://www.google.com",
		"http://www.bing.com",
		"http://www.apache.org",
		"http://www.amazon.com",
		"http://www.gmail.com",
	}

	// define the number of goroutines you intend to use
	waitGroup.Add(len(urls))
	go worker(urls)

	waitGroup.Wait()
	fmt.Println("Done!\n")
}

func worker(urls []string) {
	for _, v := range urls {
		go fetchUrl(v)
	}
}

func fetchUrl(url string) {
	// signal when the goroutine is done.
	defer waitGroup.Done()

	// make the call and collect the response
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching the URL: %s", err)
	} else {
		defer response.Body.Close()
		byte_array, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading the response: %s", err)
		} else {
			fmt.Println(string(byte_array))
		}
	}
}
