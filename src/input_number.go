// begin with the main package
package main

// import fmt for formatted IO,
// os for reading arguments to the program
// strconv to translate string parameter to int 
import ("fmt"; "os"; "strconv")


func main() {
  // if the length of arguments to the program is
  // not two (program name and the number), then
  // print out a message of how to use the program and exit

  if len(os.Args) != 2 { // note how to check for arguments
    fmt.Println("Usage: go run input_number.go <number>")
    os.Exit(1) // note how to exit the program
  }

  // convert the string (argument passed) into a int
  // using the Atoi function from strconv package
  // this will convert a string to an int. Itoa will
  // convert a number into string. 
  input, err := strconv.Atoi(os.Args[1]) // note how there are 2 return parameters from the function

  // if the conversion is successful, err returned will be nil (null in Java)
  if err != nil {
    // In case it doesn't, then it means that conversion went through some problem
    // In that case, print the error and exit.
    fmt.Println("Incorrect number passed: ", err)
    os.Exit(1)
  } else {
    fmt.Println("conversion was successful. Err = ", err)
  }

  // else execute the business logic
  if input > 2000 {
    fmt.Println("Wise choice!")
  } else {
    fmt.Println("Hmm... Too bad!")
  }
}
