package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}

// here the inbound (src) is from Acme to joe's proxy website to which Acme is allowing to connect. The following code
// is from the point of view of joesproxy.com
func handle(src net.Conn) {
	defer src.Close() // Close ACME → proxy connection
	// so from the proxy server we connect to joescatcam which is blocked by Acme
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Println("Unable to connect to our unreachable host")
		return // Don't crash the server
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking. Here we copy the inbound data (from Acme) to joesproxy.com to
	// joescatcam
	go func() {
		// copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy our destination's output back to our source. Here we are relaying data from joescatcam back to joesproxy.com
	// which gets routed to Acme
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
