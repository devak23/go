package main

import (
	"fmt"
	"strings"
)

func FunWithStringsMain() {
	helloWorld := "\tHello, World!"
	helloMars := strings.Replace(helloWorld, "World", "Mars", 1)
	fmt.Println("Length of the string:", len(helloMars), helloMars)
	helloMars = strings.TrimSpace(helloMars)
	fmt.Println("Length of the string:", len(helloMars), helloMars)
	fmt.Println("substring: ", helloMars[0:6])
	fmt.Println("starts with 'Hello, Mars'?", strings.HasPrefix(helloMars, "Hello"))
	fmt.Println("ContainsAny check: contains Mars? ", strings.ContainsAny(helloMars, "mars"))
	fmt.Println("Contains check: contains Mars? ", strings.Contains(helloMars, "mars"))
	fmt.Println("Index of 'Mars':", strings.Index(helloMars, "Mars"))
	fmt.Println("ContainsFunction (contains vowel check): contains Mars? ", strings.ContainsFunc(helloMars, hasVowel))
	fmt.Println("ToLower:", strings.ToLower(helloMars))
}

func hasVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}
