// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestSearchCards(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("search", "cards-api-example-response.json").URL
	cards, err := c.SearchCards("testQuery", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 1 {
		t.Errorf("Expected 1 card search result. Got %d.", len(cards))
	}
}

func TestSearchBoards(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("search", "boards-api-example-response.json").URL
	boards, err := c.SearchBoards("testQuery", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(boards) != 1 {
		t.Errorf("Expected 1 board search result. Got %d.", len(boards))
	}
}

func TestSearchMembers(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("search", "members-api-example-response.json").URL
	members, err := c.SearchMembers("testQuery", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 3 {
		t.Errorf("Expected 3 member search result entries. Got %d.", len(members))
	}
}
