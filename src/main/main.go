package main

import (
	gbe "gobyexample"
	"fmt"
	"assorted"
)

func main() {
	fmt.Println("============ Executing gobyexample.HelloMain =================")
	gbe.HelloMain()
	fmt.Println("============ Executing gobyexample.ArraysMain =================")
	gbe.ArraysMain()
	fmt.Println("============ Executing gobyexample.ConstantsMain =================")
	gbe.ConstantsMain()
	fmt.Println("============ Executing gobyexample.ForMain =================")
	gbe.ForMain()
	fmt.Println("============ Executing gobyexample.IfElseMain =================")
	gbe.IfelseMain()
	fmt.Println("============ Executing gobyexample.SwitchMain =================")
	gbe.SwitchMain()
	fmt.Println("============ Executing gobyexample.ValuesMain =================")
	gbe.ValuesMain()
	fmt.Println("============ Executing gobyexample.VariablesMain =================")
	gbe.VariablesMain()


	fmt.Println("============ Executing assorted.ConstantsMain =================")
	assorted.ConstantsMain()
	fmt.Println("============ Executing assorted.DoubleNumberMain =================")
	assorted.DoubleNumberMain()
	fmt.Println("============ Executing assorted.FunctionsMain =================")
	assorted.FunctionsMain()
	fmt.Println("============ Executing assorted.IntAndFloatMain =================")
	assorted.IntAndFloatMain()
	fmt.Println("============ Executing assorted.SampeStructMain =================")
	assorted.SampleStructMain()
}
