package assorted

import (
	"sync"
	"fmt"
	"net/http"
	"io/ioutil"
)

var urls = []string {
	"http://www.yahoo.com",
	"http://www.msn.com",
	"http://www.nytimes.com",
}
var waitGroup sync.WaitGroup

func FetchUrlsMain() {
	// define the number of goroutines you intend to use
	waitGroup.Add(len(urls))

	for _, v := range urls {
		go fetchUrl(v)
	}
	waitGroup.Wait()
	fmt.Println("Done!\n")
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
		//fmt.Println(response)
	}
}
