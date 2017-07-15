package assert

import (
  "testing"
)

func TestAssertion(t *testing.T) {
  a := 42
  AssertEqual(t, a, 42, "some message")
  AssertEqual(t, a, 43, "This message is displayed instead of default ones")
}
