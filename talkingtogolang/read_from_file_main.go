package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFromFileMain() {
	fmt.Println("Reading from a file...")
	file, _ := os.OpenFile("./hello.txt", os.O_RDONLY, 0644)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				fmt.Println("> ", line)
			} else {
				//log.Fatalf("Error: ", err)
				// NOTE: log.Fatal() calls os.Exit(1) which terminates the entire program immediately, rather than just
				//exiting the loop gracefully. Therefore, any subsequent function calls will not be executed.
				fmt.Println("ERROR: ", err)
			}
			break
		}
		fmt.Println("> ", line)
	}
}
