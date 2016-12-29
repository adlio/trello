// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
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

func testClient() *Client {
	c := NewClient("user", "pass")
	c.testMode = true
	return c
}
