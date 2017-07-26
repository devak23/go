package fundamentals

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile reads a file name and returns a byte array and error if any
func ReadFile(filename string) ([]byte, error) {
	// open a file for reading
	f, err := os.Open(filename)
	defer f.Close()
	// if there is an error, return nil
	if err != nil {
		fmt.Println("There was an error opening the file", err)
		return nil, err
	}

	// else read the file using ioutil.ReadAll() function
	return ioutil.ReadAll(f)
}

// FileReadMain is the main function
func FileReadMain() {
	content, err := ReadFile("/home/abhay/Workspace/go/sherlock.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
