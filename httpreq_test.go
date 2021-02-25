package tiktokgo

import (
	"testing"
)

func TestSetUserAgent(t *testing.T) {
	SetUserAgent("hello world")
	if defaultHeaders["User-Agent"] != "hello world" {
		t.Error(defaultHeaders["User-Agent"])
	}
}
