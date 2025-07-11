package main

import (
	. "fmt"
	"strconv"
)

// =====================================
// 1. Basic concepts
// =====================================

// Simple function type
type operation func(int, int) int

// Basic operations
func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}

func basicFunctionExample() {
	Println("=== Basic Function Example ===")
	// Functions can be assigned to a variable
	var op operation = add
	Println("5 + 3 = ", op(5, 3))

	// Functions can be passed as arguments to another functions
	result := calculate(10, 5, multiply)
	Println("10 * 5 = ", result)

	// Functions can be defined anonymously using lambdas
	subtract := func(a, b int) int {
		return a - b
	}
	Println("10 - 5 = ", subtract(10, 5))
}

func calculate(a, b int, op operation) int {
	return op(a, b)
}

// =====================================
// 2. Higher Order functions
// =====================================

// Function that returns a function (closure)
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Function that returns different operation based on input
func getOperations(op string) operation {
	switch op {
	case "+":
		return add
	case "*":
		return multiply
	default:
		return func(a, b int) int { return 0 }
	}
}

func higherOrderFunctionsExample() {
	Println("\n===Higher Order Functions===")

	// Closures - functions that capture variables from their environment
	add5 := makeAdder(5)
	Println("add5(10) = ", add5(10))

	// function factory
	addOp := getOperations("+")
	multOp := getOperations("*")
	Println("addOp(3,4) = ", addOp(3, 4), "\nmultOp(3,4) = ", multOp(3, 4))
}

// =========================================================
// 3. Functional Collection operations - Map, filter, reduce
// =========================================================

// Map: Transform each element in a slice
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Map function with Generics
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter - Keep elements that satisfy a condition
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce - Combine all elements into a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}

	return result
}

func functionalCollectionExample() {
	Println("\n=== Functional Collection Operations ===")
	numbers := []int{1, 2, 3, 4, 5}

	// Map: Double each number
	doubled := Map(numbers, func(x int) int { return x * 2 })
	Println("Doubled {1,2,3,4,5} = ", doubled)

	// Map: Convert to strings
	strings := Map(numbers, func(x int) string { return strconv.Itoa(x) })
	Println("As Strings = ", strings)

	// Filter to only keep even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	Println("Evens = ", evens)

	// Reduce: Sum of all numbers
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	Println("Sum = ", sum)

	// Reduce: Find max
	maximum := Reduce(numbers, 0, func(acc, x int) int {
		if x > acc {
			return x
		}
		return acc
	})
	Println("Max = ", maximum)
}

// =========================================================
// 4. Function Composition - combining functions
// =========================================================

// Compose two functions
func Compose[T, U, V any](f func(U) V, g func(T) U) func(T) V {
	return func(x T) V {
		return f(g(x))
	}
}

func Pipeline[T any](value T, operations ...func(T) T) T {
	result := value
	for _, op := range operations {
		result = op(result)
	}
	return result
}

func compositionExample() {
	Println("\n=== Composition ===")

	// Simple composition
	double := func(x int) int { return x * 2 }
	sqr := func(x int) int { return x * x }

	// Compose: first double then square
	doubleAndSquare := Compose(sqr, double)
	Println("doubleAndSquare(5) = ", doubleAndSquare(5))

	// Pipeline example
	result := Pipeline(5,
		func(x int) int { return x * 2 }, // 10
		func(x int) int { return x + 1 }, // 11
		func(x int) int { return x * 3 }, // 33
	)
	Println("Pipeline(5) = ", result)
}

// =========================================================
// 5. Partial application and currying
// =========================================================

// PartialAdd - Fix some arguments
func PartialAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func CurriedAdd(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

func Curry2[T, U, V any](fn func(T, U) V) func(T) func(U) V {
	return func(t T) func(U) V {
		return func(u U) V {
			return fn(t, u)
		}
	}
}

func partialAndCurryExample() {
	Println("\n=== Partial Application and Currying ===")
	// Partial Application
	add10 := PartialAdd(10)
	Println("add10(12) = ", add10(12))

	// Currying
	curriedAdd := CurriedAdd(1)
	addWith1And2 := curriedAdd(2)
	result := addWith1And2(3)
	Println("CurriedAdd(1)(2)(3) = ", result)

	// Generic currying
	curriedMultiply := Curry2(multiply)
	multiplyBy3 := curriedMultiply(3)
	Println("multiplyBy3(4) = ", multiplyBy3(4))
}

// =========================================================
// 6. Monadic Patterns - Optional/Maybe type
// =========================================================

// Optional type to handle nil values functionality
type Optional[T any] struct {
	value *T
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

func None[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

func (o Optional[T]) IsSome() bool {
	return o.value != nil
}

func (o Optional[T]) IsNone() bool {
	return o.value == nil
}

func (o Optional[T]) Unwrap() T {
	if o.value == nil {
		panic("Cannot unwrap None")
	}
	return *o.value
}

func (o Optional[T]) UnwrapOr(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}

// Map over Optional (functor)
func (o Optional[T]) Map(fn func(T) T) Optional[T] {
	if o.value == nil {
		return None[T]()
	}
	result := fn(*o.value)
	return Some(result)
}

// FlatMap for Optional (monad)
func (o Optional[T]) FlatMap(fn func(T) Optional[T]) Optional[T] {
	if o.value == nil {
		return None[T]()
	}
	return fn(*o.value)
}

// Safe division that returns Optional
func safeDivide(a, b int) Optional[int] {
	if b == 0 {
		return None[int]()
	} else {
		return Some(a / b)
	}
}

func monadicExamples() {
	Println("\n=== Monadic Patterns (Optional) ===")

	// Basic Optional use
	some := Some(42)
	none := None[int]()
	Println("Some(42) = ", some.UnwrapOr(0))
	Println("None = ", none.UnwrapOr(0))

	// Map over Optional
	doubled := some.Map(func(x int) int { return x * 2 })
	Println("Some(42) mapped with * 2 = ", doubled.UnwrapOr(0))

	// Chaining operations with Flatmap
	result := Some(10).FlatMap(func(x int) Optional[int] {
		return safeDivide(x, 2)
	}).Map(func(x int) int {
		return x + 1
	})
	Println("Some(10) -> safeDivide(2) -> add(1) = ", result.UnwrapOr(0))

	// Safe division chain
	divResult := Some(100).
		FlatMap(func(x int) Optional[int] { return safeDivide(x, 10) }).
		FlatMap(func(x int) Optional[int] { return safeDivide(x, 2) })
	Println("100 / 10 / 2 = ", divResult.UnwrapOr(0))
}

func main() {
	Println("FUNCTIONAL PROGRAMMING IN GO - EXAMPLES")
	Println("========================================")
	basicFunctionExample()
	higherOrderFunctionsExample()
	functionalCollectionExample()
	compositionExample()
	partialAndCurryExample()
	monadicExamples()
	complexFunctionalExample()
}
