package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

// The name of a test function must begin with the word Test, or Go won’t recognise it as a test. It also has to accept
// a parameter whose type is *testing.T. testing.T is a struct that serves as the core for writing and controlling tests.

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

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{2, 2, 4},
		{-2, -2, -4},
		{0, 0, 0},
		{2.5, 2.5, 5},
		{-2.4, 0.4, -2},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		assertOutcome(t, tc.want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{4, 2, 2},
		{0, 0, 0},
		{2.5, 2.5, 0},
		{-2.4, 0.4, -2.8},
		{2, 4, -2},
		{4, 2, 2},
		{0, 0, 0},
		{-2.5, -2.5, 0},
		{-2.4, 0.4, -2.8},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		assertOutcome(t, tc.want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{2, 2, 4},
		{-2, -2, 4},
		{0, 0, 0},
		{2.5, 2.5, 6.25},
		{-2.4, 0.4, -0.96},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		assertOutcome(t, tc.want, got)
	}
}

func TestDivideWithValidInputs(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{4, 2, 2},
		{2.5, 2.5, 1},
		{-2.4, 0.4, -6},
		{2, 4, 0.5},
		{4, 2, 2},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Wanted no errors for valid inputs, but got: %v", err)
		}
		assertOutcome(t, tc.want, got)
	}
}

func TestDivideWithInvalidInputs(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{0, 0, 0},
		{2, 0, 0},
	}

	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)
		if err == nil {
			t.Fatalf("Wanted an error for invalid inputs, but got: %v", err)
		}
	}
}

func assertOutcome(t *testing.T, want float64, got float64) {
	if !almostEqual(want, got) {
		t.Errorf("wanted %f, got %f", want, got)
	}
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-10 // or 0.000,000,000,1
	return math.Abs(a-b) < epsilon
}
