package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			fmt.Printf("trying %s\n", address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				return
			}
			_ = conn.Close()
			fmt.Printf("%d is open.\n", j)
		}(i)
		wg.Wait()
	}
}

// This version of the program is better, but still incorrect. If you run this multiple times against multiple hosts, you
// might see inconsistent results. Scanning an excessive number of hosts or ports simultaneously may cause network or
// system limitations to skew your results
