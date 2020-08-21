package template

import (
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	handler := NewHTTPHandler()
	handler.Handle("http://example.com/abc.zip")
}

func TestFtpHandler(t *testing.T) {
	handler := NewFtpHandler()
	handler.Handle("ftp://example.com/abc.zip")
}
