package main

import fm "fmt" // creating an alias called fm

const Pi = 3.14159
// constants can overflow if they are assigned numeric variables with too little
// precision. If that occurs, its a compile time error. Multiple assignments are 
// allowed as follows
const beef, two, c  = "meat", 2, "veg"

const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday = 1,2,3,4,5,6,7

const (
	blue, yellow, red = 3, 2, 1
	orange, pink, purple = 6, 7, 8
)

// you can use 'iota' to provide sequential values as in the following case
const (
	JAN = iota
	FEB
	MAR
	APR
	MAY
	JUN
	JUL
	AUG
	SEP
	OCT
	NOV
	DEC
)

type Color int
const (
	RED Color = iota;
	ORANGE; YELLOW; BLUE; GREEN; INDIGO; VIOLET;
)

func main() {
	fm.Println("This is printed using alias")
	print("ABC")
	println("\nTHis is on a new line")
	fm.Println("Καλημέρα κόσμε; or こんにちは 世界\n")
	fm.Println("value of Pi=", Pi)

	fm.Println("constants have the values: ", beef, two, c)

	fm.Println("sunday = ", Sunday)
	fm.Println("purple = ", purple)

	fm.Println("AUG = ", AUG)
	for i := RED; i <= VIOLET; i++ {
		fm.Println(i)
	}
}
