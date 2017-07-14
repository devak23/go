package gobyexample

import "fmt"
import "time"

func SwitchMain() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// you can use commas to separate multiple expressions
	// in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is a weekend")
	default:
		fmt.Println("Today is a weekday.", "[Today is", time.Now().Weekday(), "]")
	}

	// switch without an expression is another way of expressing
	// if-else logic
	t := time.Now()
	fmt.Println("Now = ", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// a type switch compares types and not values
	// you can use this to discover the type of an
	// interface value. In this case, the variable t
	// will have the type corresponding to its value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Dont know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
