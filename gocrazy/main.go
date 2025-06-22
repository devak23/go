package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Called before calling main!")
}

func main() {
	programName, questions := os.Args[0], os.Args[1:]
	fmt.Println("ProgramName, questions = ", programName, questions)

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
	LoadingEnvMain()
	GoRoutineMain()
	GoChannelMain()
	ReadFromChannelMain()
	ReadFromBufferedChannelMain()
	WorkingWithDiffSourcesMain()
}
