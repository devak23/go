package algorithms

import "testing"

// use values that you know are right
var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

// TestFibFunc tests the FibFunc function exposed
func TestFibFunc(t *testing.T) {
	fn := FibFunc()
	for i, v := range tests {
		if val := fn(); val != v {
			t.Fatalf("at index %d, expected %d, got %d", i, val, v)
		}
	}
}
