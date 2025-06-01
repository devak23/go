package assorted

import (
	"fmt"
	"unicode/utf8"
)

func StringLengthMain() {
	mystring := "Hello Go"
	length(mystring)

	mystring = "你好"
	length(mystring)

	mystring = "こんにちは"
	length(mystring)
}

func length(mystring string) {
	fmt.Printf("length of \"%s\" in bytes = %d\n", mystring, len(mystring))
	fmt.Printf("length of \"%s\" in runes = %d\n", mystring, utf8.RuneCountInString(mystring))
}
