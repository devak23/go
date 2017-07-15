package goinpractise

import (
  "testing"
  "assert"
)

func TestMultipleReturns(t *testing.T) {
  name1, name2 := MultipleReturns()
  assert.AssertEquals(t, name1, "Abhay", "Matches")
  assert.AssertEquals(t, name2, "Soham", "Matches")
}

func TestMultipleReturnsByPassingParams(t *testing.T) {
  ret1, ret2 := MultipleReturnsWithParams("Manik", "Soham")
  assert.AssertEquals(t, ret1, "Manik", "")
  assert.AssertEquals(t, ret2, "Soham", "")
}
