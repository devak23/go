package goinpractise

import (
	"assert"
	"testing"
)

func TestReadStatusFromGoogleAtPort80(t *testing.T) {
	url := "www.google.com:80"
	response := ConnectAndGet(url)
	assert.AssertNotNil(t, response, "")
}

func TestReadStatusFromGoLangAtPort80(t *testing.T) {
	url := "golang.org:80"
	response := ConnectAndGet(url)
	assert.AssertNotNil(t, response, "")
}

func TestReadStatusFromApacheAtPort80(t *testing.T) {
	url := "apache.org:80" //had forgotten the port earlier
	response := ConnectAndGet(url)
	assert.AssertNotNil(t, response, "")
	assert.AssertEquals(t, response, "HTTP/1.1 200 OK\r\n", "")
}
