package assert

import (
	"fmt"
	"testing"
)

func AssertEquals(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func AssertNotNil(t *testing.T, a interface{}, message string) {
	if a == nil {
		t.Fatal("was not expecting a null input for assertion")
	}
}
