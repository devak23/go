package main

import (
	"fmt"
	"os"
)

func main() {
	DeferringMain()
	SortingMain()
	TypesAgainDemo()
	ReadStructFromFileWithDecoderMain()
	ReadStructFromFileMain()
	WriteStructIntoFileMain()
	MarshallStructMain()
	UsingStructMain()
	ReadFromFileMain()
	//ReadFromTerminalMain()
	programName, questions := os.Args[0], os.Args[1:]
	fmt.Println("ProgramName, questions = ", programName, questions)
	LoadingEnvMain()
	GoRoutineMain()
	GoChannelMain()
}
