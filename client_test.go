package trello

import (
	"testing"
)

func TestGetWithBadURL(t *testing.T) {
	c := testClient()
	target := map[string]interface{}{}
	c.BaseURL = "gopher://test"
	err := c.Get("members", Defaults, &target)
	if err == nil {
		t.Fatal("Get() should fali with a bad URL")
	}
}

func testClient() *Client {
	c := NewClient("user", "pass")
	return c
}
