// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
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
	c := NewClient("", "")
	c.testMode = true
	c.BaseURL = "http://synthetic.invalid"
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
	newC.Client = &http.Client{
		Transport: &mockTransport{
			RoundTripFunc: func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Context() != ctx {
					t.Fatal("request should use the new context")
				}
				return syntheticResponse(req)
			},
		},
	}
	target := map[string]interface{}{}
	if err := newC.Get("members", Defaults(), &target); err != nil {
		t.Fatal(err)
	}

	if calls != 1 {
		t.Fatal("Get() should have used the mocked transport")
	}
}

func TestHTTPMethodsHonorClientContext(t *testing.T) {
	methods := []struct {
		name string
		call func(*Client) error
	}{
		{name: "GET", call: func(c *Client) error {
			target := map[string]interface{}{}
			return c.Get("synthetic", Defaults(), &target)
		}},
		{name: "PUT", call: func(c *Client) error {
			target := map[string]interface{}{}
			return c.Put("synthetic", Defaults(), &target)
		}},
		{name: "POST", call: func(c *Client) error {
			target := map[string]interface{}{}
			return c.Post("synthetic", Defaults(), &target)
		}},
		{name: "POST with body", call: func(c *Client) error {
			target := map[string]interface{}{}
			return c.PostWithBody("synthetic", Defaults(), &target, "synthetic.txt", strings.NewReader("synthetic body"))
		}},
		{name: "DELETE", call: func(c *Client) error {
			target := map[string]interface{}{}
			return c.Delete("synthetic", Defaults(), &target)
		}},
	}

	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			c := NewClient("", "")
			c.testMode = true
			c.BaseURL = "http://synthetic.invalid"
			c = c.WithContext(ctx)
			c.Client = &http.Client{
				Transport: &mockTransport{RoundTripFunc: syntheticResponse},
			}

			if err := method.call(c); !errors.Is(err, context.Canceled) {
				t.Fatalf("expected request to honor the canceled client context, got %v", err)
			}
		})
	}
}

func syntheticResponse(req *http.Request) (*http.Response, error) {
	if err := req.Context().Err(); err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("{}")),
		Request:    req,
	}, nil
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
