package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Printf("trying: %s\n", address)
		conn, err := net.Dial("tcp", address) // Dial(network, <address string>)
		if err != nil {
			// port is closed or filtered
			continue
		}
		_ = conn.Close()
		fmt.Printf("%d is open.\n", i)
	}
}

// scanme.nmap.org is a free service provided by Fyodor, the creator of Nmap. When we’re scanning, we will be polite.
// He requests, "Try not to hammer on the server too hard. A few scans in a day is fine, but don’t scan 100 times a day."
