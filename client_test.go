// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"context"
	"net/http"
	"testing"
)

func TestGetWithBadURL(t *testing.T) {
	c := testClient()
	target := map[string]interface{}{}
	c.BaseURL = "gopher://test"
	err := c.Get("members", Defaults(), &target)
	if err == nil {
		t.Fatal("Get() should fail with a bad URL")
	}
}

func TestWithContext(t *testing.T) {
	c := testClient()
	if c.ctx != context.Background() {
		t.Fatal("NewClient() should use context.Background()")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newC := c.WithContext(ctx)
	if newC == c {
		t.Fatal("WithContext() should return a new client")
	}

	if newC.ctx != ctx {
		t.Fatal("WithContext() should return a client with the given context")
	}

	var calls int
	mt := &mockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			calls++
			if req.Context() != ctx {
				t.Fatal("Get() should be using the new context")
			}
			return http.DefaultTransport.RoundTrip(req)
		},
	}
	newC.Client = &http.Client{
		Transport: mt,
	}
	newC.Get("members", nil, nil)

	if calls != 1 {
		t.Fatal("Get() should have used the mocked transport")
	}
}

type mockTransport struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFunc(req)
}

func testClient() *Client {
	c := NewClient("user", "pass")
	c.testMode = true
	return c
}
