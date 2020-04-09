// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestDefaultArguments(t *testing.T) {
	args := Defaults()
	queryString := args.ToURLValues().Encode()
	if queryString != "" {
		t.Errorf("Query string should be blank for default Trello arguments. Got '%s' instead.", queryString)
	}
}

func TestSingleArgument(t *testing.T) {
	args := Arguments{"limit": "1000"}
	queryString := args.ToURLValues().Encode()
	if queryString != "limit=1000" {
		t.Errorf("Expected 'limit=1000', but got '%s' instead.", queryString)
	}
}
