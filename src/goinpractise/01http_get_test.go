package goinpractise

import (
	"assert"
	"testing"
)

func TestHttpGetWithGoogleDotCom(t *testing.T) {
	url := "http://www.google.com"
	response, error := HttpGet(url)
	assert.AssertEquals(t, error, nil, "")
	assert.AssertNotNil(t, response, "")
}

func TestHttpGetWithJunkWebsiteGivesError(t *testing.T) {
	url := "http://hhhhhhh:8080/"
	response, error := HttpGet(url)
	assert.AssertEquals(t, response, "", "")
	assert.AssertNotNil(t, error, "")
}
