package waytogo

import (
	"fmt"
	"bufio"
	"os"
)

func WordLetterCountMain() {
	fmt.Printf("Enter a line of text (end with letter S): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('S')
	if err != nil {
		fmt.Println("Erroneous input. Program terminated")
		os.Exit(1)
	}
	charCount, wordCount, lineCount := analyzeLine(input)
	fmt.Printf("\nNumber of characters = %d\nNumber of lines = %d\nNumber of words = %d\n", charCount, lineCount, wordCount)
}

func analyzeLine(line string) (int, int, int) {
	charCount := len(line)
	lineCount, wordCount := 1,0
	for index, _ := range line {
		if line[index] == '\n' {
			if index > 0 && line[index-1] != ' ' {
				wordCount++
			}
			lineCount++
		} else if (index > 0 && line[index] != ' ' && line[index-1] != ' ') || (index == 0 && line[0] != ' ') {
			wordCount++
		}
	}

	return charCount, wordCount, lineCount
}
