package main

// Based on Go’s documentation for the data type, Conn implements the Read([]byte) and Write([]byte) functions as defined
// for the Reader and Writer interfaces. Therefore, Conn is both a Reader and a Writer. This makes sense logically, as
// TCP connections are bidirectional and can be used to send (write) or receive (read) data.
// After creating an instance of Conn, we’ll be able to send and receive data over a TCP socket. However, a TCP server
// can’t simply manufacture a connection; a client must establish a connection. So we can use
// net.Listen(network, address string) to first open a TCP listener on a specific port. Once a client connects, the
// Accept() method creates and returns a Conn object that we can use for receiving and sending data.

import (
	"io"
	"log"
	"net"
)

func main() {
	// Setting up the listener by binding the TCP port to 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	// Next, an infinite loop y ensures that the server will continue to listen for connections even after one has been
	// received
	for {
		// Wait for the connection. The call listener.Accept() blocks execution as it awaits client connections.
		// When a client connects, this function returns a Conn instance.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// ... which is then passed into the echo function (using goroutine for concurrency)
		go echo(conn)
	}
}

// echo is a handler function that simply echoes received data. Conn is both a Reader and a Writer
// (it implements the Read([]byte) and Write([]byte) interface methods)
func echo(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	// create a buffer to store received data
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into the buffer
		size, err := conn.Read(b[0:])

		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error: ", err)
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b[0:size]))

		// Send data via conn.Write
		log.Println("Writing back data...")
		response := append([]byte("echo:"), b[:size]...)
		if _, err := conn.Write(response); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
