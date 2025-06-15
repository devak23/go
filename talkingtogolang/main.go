package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	readFromTerminal()
	readFromFile()
}

func readFromTerminal() (*bufio.Reader, string, error) {
	fmt.Println("Reading from a terminal ...")
	fmt.Print("What is your name? - ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s!\n", line)
	return reader, line, err
}

func readFromFile() {
	fmt.Println("Reading from file ...")
	file, _ := os.OpenFile("./hello.txt", os.O_RDONLY, 0666)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					fmt.Printf("> %s\n", line)
				}
			}
			log.Fatal(err)
			return
		}
		fmt.Printf("> %s\n", line)
	}
}
