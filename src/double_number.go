// All go programs must have a main package where the execution begins
package main

// import the formated IO package for print statements
import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  // use the Scanf and pass the address of the input variable
  fmt.Scanf("%f", &input)

  // declare and initialize the output variable by doubling the input
  output := input * 2

  // print the output variable
  fmt.Println("double = ", output)
}
