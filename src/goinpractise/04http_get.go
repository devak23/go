package goinpractise

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HttpGet connects to the url specified and performs a GET
// request to return a response
func HttpGet(url string) (string, error) {
	// open connection to a given url
	response, errConn := http.Get(url)
	if errConn != nil {
		fmt.Println(errConn)
		return "", errConn
	}

	// read the response body
	body, errorBody := ioutil.ReadAll(response.Body)

	returnValue := ""
	if errorBody == nil {
		// collect the response into a string object
		returnValue = string(body)
		// close the response
		response.Body.Close()
	}

	// return the response
	return returnValue, errorBody
}
