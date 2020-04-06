// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestGetMembersOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("members", "board-members-api-example.json").URL
	members, err := board.GetMembers(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 3 {
		t.Errorf("Expected 3 members, got %d", len(members))
	}
}

func TestGetMembersInOrganization(t *testing.T) {
	organization := testOrganization(t)
	organization.client.BaseURL = mockResponse("members", "trelloapps.json").URL
	members, err := organization.GetMembers(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 7 {
		t.Errorf("Expected 3 members, got %d", len(members))
	}
}

func TestGetMembersOnCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("members", "card-members-api-example.json").URL
	members, err := card.GetMembers(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 1 {
		t.Errorf("Expected 1 member, got %d", len(members))
	}
}
