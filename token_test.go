// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	token := testToken(t)

	if token.Identifier != "Name of Application" {
		t.Errorf("Expected 'Name of Application' blah, got '%s'.", token.Identifier)
	}

	if token.DateExpires != nil {
		t.Error("token.DateExpires should have been nil for this token.")
	}

	if len(token.Permissions) != 3 {
		t.Errorf("Expected 3 permissions, got %d.", len(token.Permissions))
	}

	boardPerm := token.Permissions[1]
	if boardPerm.IDModel != "*" || boardPerm.Read != true {
		t.Error("Expected read permissions on all boards.")
	}

}

func TestGetExpiringToken(t *testing.T) {
	client := testClient()
	client.BaseURL = mockResponse("tokens", "token-expiring.json").URL
	token, err := client.GetToken("tOkenId", Defaults())
	if err != nil {
		t.Error(err)
	}

	if token.DateExpires == nil {
		t.Error("Token should have an expiration date.")
	}

	if token.DateExpires.Format(time.Kitchen) != "4:25AM" {
		t.Errorf("Expected 4:25AM expiration time. Got %s.", token.DateExpires.Format(time.Kitchen))
	}
}

func testToken(t *testing.T) *Token {
	client := testClient()
	client.BaseURL = mockResponse("tokens", "token.json").URL
	token, err := client.GetToken("tOkenId", Defaults())
	if err != nil {
		t.Error(err)
	}
	return token
}
