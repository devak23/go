package assorted

import (
	"bytes"
	"fmt"
)

// This is the main function that gets called from main.go
func EfficientStringConcatMain() {
	names := []string {"Soham, Abhay, Manik"}

	joined := join(names, ", ")

	fmt.Printf("Concatenated string: %s\n", joined)

}

func join(strings []string, delimit string) string {
	var buffer bytes.Buffer

	count := 0
	for _, s := range strings {
		count++
		// writes the string into the buffer. The buffer will grow as more
		// content gets written to it.
		buffer.WriteString(s)

		// this logic is to remove the trailing comma
		if count < len(strings) {
			buffer.WriteString(delimit)
		}
	}

	return buffer.String()
}