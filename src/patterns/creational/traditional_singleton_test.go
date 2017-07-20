package creational

import (
	"testing"
	"assert"
)

func TestGetElvis(t *testing.T) {
	counter1 := GetElvis()

	if counter1 == nil {
		t.Error("expected pointer to get a ElvisSingleton after calling GetElvis(), but got", nil)
	}

	expectedCounter := counter1
	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		assert.AssertEquals(t, expectedCounter, currentCounter, "" )
	}
}
