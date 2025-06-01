package assert

import (
	"testing"
)

func TestAssertion(t *testing.T) {
	a := 42
	AssertEquals(t, a, 42, "some message")
	AssertEquals(t, a, 43, "This message is displayed instead of default ones")
}
