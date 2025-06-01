package calculator_test

import (
	"calculator"
	"testing"
)

// The name of a test function must begin with the word Test, or Go won’t recognise it as a test. It also has to accept
//a parameter whose type is *testing.T. testing.T is a struct that serves as the core for writing and controlling tests.

// Every Go test function must accept exactly one parameter of type *testing.T. This is a requirement of Go to recognize
// a function as a test. The parameter name t is just a convention. You can name it differently, but t is a convention.
// The t (*testing.T) provides a context which has various methods to report results such as
// t.Error(), t.Fail(), t.Fatal(), t.Fatalf() for testing failures
// t.Skip(), t.Parallel(), t.Run() for controlling test execution
// t.Log(), t.Logf() for logging information
// t.Name(), t.Failed() for accessing test metadata

// The reason it's a pointer to the struct is that it mutates the internal state of the struct. Passing an entire
// struct would not have been very efficient, so passing a 'handle' is much better!

// t.Parallel() statement is a standard prelude to tests: it tells Go to run this test concurrently with other tests,
//which saves time.

func TestAddWithInts(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	assertOutcome(t, want, got)
}

func TestAddWithFloats(t *testing.T) {
	t.Parallel()
	want := 5.0
	got := calculator.Add(3.0, 2.0)
	assertOutcome(t, want, got)
}

func TestAddWithNegativeNumbers(t *testing.T) {
	t.Parallel()
	want := -5.0
	got := calculator.Add(-3.0, -2.0)
	assertOutcome(t, want, got)
}

func TestSubtractWithInts(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	assertOutcome(t, want, got)
}

func TestSubtractWithFloats(t *testing.T) {
	t.Parallel()
	want := 4.0
	got := calculator.Subtract(6.0, 2.0)
	assertOutcome(t, want, got)
}

func TestSubtractWithNegativeNumbers(t *testing.T) {
	t.Parallel()
	want := -4.0
	got := calculator.Subtract(-6.0, -2.0)
	assertOutcome(t, want, got)
}

func TestSubtractWithDecimals(t *testing.T) {
	t.Parallel()
	want := 0.25
	got := calculator.Subtract(1.0, 0.75)
	assertOutcome(t, want, got)
}

func TestAddWithDecimals(t *testing.T) {
	t.Parallel()
	want := 0.75
	got := calculator.Add(1.0, -0.25)
	assertOutcome(t, want, got)
}

func TestMultiplyWithPositiveNumbers(t *testing.T) {
	t.Parallel()
	var want float64 = 25
	got := calculator.Multiply(5, 5)
	assertOutcome(t, want, got)
}

func TestMultiplyWithNegativeNumbers(t *testing.T) {
	t.Parallel()
	var want float64 = -30
	got := calculator.Multiply(5, -6)
	assertOutcome(t, want, got)
}

func TestMultiplyTwoNegativeNumbersShouldLeadToPositiveResult(t *testing.T) {
	t.Parallel()
	var want float64 = 30
	got := calculator.Multiply(-5, -6)
	assertOutcome(t, want, got)
}

func assertOutcome(t *testing.T, want float64, got float64) {
	if want != got {
		t.Errorf("wanted %f, got %f", want, got)
	}
}
