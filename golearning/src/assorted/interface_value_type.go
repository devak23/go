// How do we know what specific object is being held by the interface pointer?
package assorted

import (
	"fmt"
	"strconv"
)

// define an object
type Human struct {
	name string
	age  int
}

// define an empty interface (which is implemented by everyone!)
type Element interface{}

// define a custom type called List which is nothing but an array of
// objects that implement the Element interface
type List []Element

// define a toString() sorta method for the object
func (h Human) String() string {
	return "(name: " + h.name + "- age: " + strconv.Itoa(h.age) + " years)"
}

// the main function
func InterfaceValueTypeMain() {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Human{"Dennis", 37}

	determineTypeUsingIf(list)
	determineTypeUsingSwitch(list)
}

func determineTypeUsingIf(list List) {
	fmt.Println("---------------USING IF-------------")
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and it's value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and it's value is %s\n", index, value)
		} else if value, ok := element.(Human); ok {
			fmt.Printf("list[%d] is a Human object and it's value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is a weird one!", index)
		}
	}
}

func determineTypeUsingSwitch(list List) {
	fmt.Println("---------------USING SWITCH-------------")
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and it's value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and it's value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is a Human object and it's value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is a weird one!", index)
		}
	}
}
