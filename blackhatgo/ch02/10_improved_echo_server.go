package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":20081")
	if err != nil {
		log.Fatalln("Unable to bind to the port")
	}
	log.Println("Listening on 0.0.0.0:20081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection from the client")
		}
		go improvedEcho(conn)
	}
}

func improvedEcho(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		log.Printf("Read %d bytes: %s", len(s), s)

		log.Printf("Writing data")
		if _, err := writer.WriteString("echo> " + s); err != nil {
			log.Fatalln("Unable to write data")
		}
		_ = writer.Flush()
	}
}
