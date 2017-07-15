package goinpractise

import (
	"bufio"
	"fmt"
	"net"
)

// ConnectAndGet makes a connection to a website and returns a response
// @param: url - the url of the website
// @return: returns a response of the website
func ConnectAndGet(url string) string {
	// make a connection over TCP
	conn, _ := net.Dial("tcp", url)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	return status
}

// ReadStatusMain is the public main function that can be invoked from main.go
func ReadStatusMain() {
	response := ConnectAndGet("apache.org:80")
	fmt.Printf("Response from the server:%s\n", response)
}
