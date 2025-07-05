package episode0

import "testing"

func TestBusinessLogic(t *testing.T) {
	ht := NewInMemoryDb()
	// invoke business logic
	BusinessLogic(ht)
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatal(err)
	}
	if string(val) != "world" {
		t.Fatalf("expected %s but got %s", "world", val)
	}
}

// Notes:
// 1. The file name has to end with _test
// 2. The test function has to start with Test<WhateverName>
