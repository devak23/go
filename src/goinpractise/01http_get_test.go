package goinpractise

import (
	"assert"
	"testing"
)

func TestHttpGetWithGoogleDotCom(t *testing.T) {
	url := "http://www.google.com"
	response, err := HttpGet(url)
	assert.AssertEquals(t, err, nil, "")
	assert.AssertNotNil(t, response, "")
}

func TestHttpGetWithJunkWebsiteGivesError(t *testing.T) {
	url := "http://hhhhhhh:8080/"
	response, err := HttpGet(url)
	assert.AssertEquals(t, response, "", "")
	assert.AssertNotNil(t, err, "")
}
