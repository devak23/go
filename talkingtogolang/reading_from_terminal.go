package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromTerminal() {
	fmt.Println("Reading from a terminal ...")
	fmt.Print("What is your name? - ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s!\n", strings.TrimSpace(line))
}

/**
 * reader.ReadString('\n') is designed to read up to and including the delimiter character. This is by design because:
 * It preserves the complete input stream - the delimiter is part of the data that was actually present
 * It allows you to distinguish between different types of line endings (Unix \n, Windows \r\n, old Mac \r)
 * It's consistent with how most buffered readers work across programming languages
 *
 * But since it reads the delimiter as well, the carriage return and line feed (or just newline) also gets added to the
 * the input. Therefore to remove that extra newline, we use strings.TrimSpace() to trim the whitespace from the end
 * of the line.
 */
