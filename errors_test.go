package trello

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRateLimitError(t *testing.T) {
	rc := ioutil.NopCloser(&bytes.Buffer{})
	resp := &http.Response{
		Body:       rc,
		StatusCode: http.StatusTooManyRequests,
	}
	e := makeHTTPClientError("/url/string", resp)
	if !IsRateLimit(e) {
		t.Error("Expected rate limit error")
	}
	if !strings.HasPrefix(e.Error(), "HTTP request failure") {
		t.Errorf("Expected error message 'HTTP request failure...', got: '%s'", e.Error())
	}
}

func TestNotFoundError(t *testing.T) {
	rc := ioutil.NopCloser(&bytes.Buffer{})
	resp := &http.Response{
		Body:       rc,
		StatusCode: http.StatusNotFound,
	}
	e := makeHTTPClientError("/url/string", resp)
	if !IsNotFound(e) {
		t.Error("Expected not found error")
	}
	if !strings.HasPrefix(e.Error(), "HTTP request failure") {
		t.Errorf("Expected error message 'HTTP request failure...', got: '%s'", e.Error())
	}
}

func TestPermissionDeniedError(t *testing.T) {
	rc := ioutil.NopCloser(&bytes.Buffer{})
	resp := &http.Response{
		Body:       rc,
		StatusCode: http.StatusUnauthorized,
	}
	e := makeHTTPClientError("/url/string", resp)
	if !IsPermissionDenied(e) {
		t.Error("Expected not found error")
	}
	if !strings.HasPrefix(e.Error(), "HTTP request failure") {
		t.Errorf("Expected error message 'HTTP request failure...', got: '%s'", e.Error())
	}
}
