// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"testing"
	"time"
)

func TestListCreatedAt(t *testing.T) {
	l := List{ID: "4d5ea62fd76aa1136000000c"}
	ts := l.CreatedAt()
	if ts.IsZero() {
		t.Error("Time shouldn't be zero.")
	}
	if ts.Unix() != 1298048559 {
		t.Errorf("Incorrect CreatedAt() time: '%s'.", ts.Format(time.RFC3339))
	}
}

func TestGetList(t *testing.T) {
	list := testList(t)
	if list.Name != "To Do Soon" {
		t.Errorf("Title incorrect. Got '%s'", list.Name)
	}
}

func TestGetListsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("lists", "board-lists-api-example.json").URL
	lists, err := board.GetLists(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(lists) != 3 {
		t.Errorf("Expected 1 list, got %d", len(lists))
	}
}

// Utility function to get the standard case Client.GetList() response
//
func testList(t *testing.T) *List {
	c := testClient()
	c.BaseURL = mockResponse("lists", "list-api-example.json").URL
	list, err := c.GetList("4eea4ff", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return list
}
