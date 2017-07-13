package assorted

import "fmt"

func add(num1 int, num2 int) int {
	var result int
	result = num1 + num2
	return result
}

func main() {
	// adding two numbers
	fmt.Println("adding 4 and 5 gives", add(4, 5))
}
