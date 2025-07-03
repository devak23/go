package main

import . "fmt"

// Function currying is often mistaken for partial functions. It is actually transforming a function that takes a single
// argument to a sequence of functions where each function takes exactly one argument.
// Func(a,b,c) int {} first takes F(a) and returns a new function Fa which takes b as the input resulting into Fa(b) giving
// Fab which takes c as input giving Fabc and returns int as output
// func F(a) 	: Fa(b)
// func Fa(b) 	: Fab(c)
// func Fabc(c)	: int
// This is done via the concept of first class functions and higher order functions. What we achieve from this is
// function composability. You can think of it as partial functions applied to a single argument.

func threeSum(a, b, c int) int {
	return a + b + c
}

// with currying, it would be
func threeSumCurried(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func main() {
	Println("Traditional: ", threeSum(10, 20, 30))
	Println("Curried: ", threeSumCurried(10)(20)(30))
}

// This goes to say that curried functions are difficult to comprehend and therefore should be used where it makes sense
// For this simple example, it doesn't make sense at all however it makes sense when you have partial application to create
// flexible functions.

// DogSpawnerCurry is a
// 1. function that takes breed as input and returns a -
// 2. function that takes Gender as input and returns a -
// 3. function that takes a Name as input and returns a Dog
func DogSpawnerCurry(breed Breed) func(Gender) NameToDogFunc {
	return func(gender Gender) NameToDogFunc {
		return func(name Name) Dog {
			return Dog{
				Name:   name,
				Breed:  breed,
				Gender: gender,
			}
		}
	}
}

// We can now see how type-aliasing really makes it very clear as to what our inputs will be
