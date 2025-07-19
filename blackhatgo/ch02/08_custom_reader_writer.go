package main

import (
	"fmt"
	"log"
	"os"
)

type CustomReader struct{}

type CustomWriter struct{}

// Reads data from stdin
func (cr *CustomReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

func (cr *CustomWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// Instantiate reader and writer
	var (
		reader CustomReader
		writer CustomWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use the reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use the writer to write output
	s, err = writer.Write(input[:s]) // it's sensible to only write the data held in the buffer and not the whole buffer itself.
	if err != nil {
		log.Fatalln("Unable to write")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
