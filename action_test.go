// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.
package trello

import (
	"testing"
)

func TestGetActionsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("actions", "board-actions-api-example.json").URL
	actions, err := board.GetActions(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(actions) != 4 {
		t.Errorf("Expected 4 actions, got %d", len(actions))
	}
}

func TestGetActionsOnList(t *testing.T) {
	list := testList(t)
	list.client.BaseURL = mockResponse("actions", "list-actions-api-example.json").URL
	actions, err := list.GetActions(Defaults)
	if err != nil {
		t.Fatal(err)
	}
	if len(actions) != 2 {
		t.Errorf("Expected 2 actions, got %d", len(actions))
	}
}

func TestGetActionsOnCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-api-example.json").URL
	actions, err := card.GetActions(Defaults)
	if err != nil {
		t.Fatal(err)
	}
	if len(actions) != 50 {
		t.Errorf("Expected 50 actions, got %d", len(actions))
	}
}
