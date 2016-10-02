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
	actions, err := board.GetActions(Defaults())
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
	actions, err := list.GetActions(Defaults())
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
	actions, err := card.GetActions(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(actions) != 50 {
		t.Errorf("Expected 50 actions, got %d", len(actions))
	}
}

func TestListAfterActionOnUpdateCard(t *testing.T) {
	a := &Action{
		Type: "updateCard",
		Data: &ActionData{
			ListBefore: &List{Name: "Before"},
			ListAfter:  &List{Name: "After"},
		},
	}
	l := ListAfterAction(a)
	if l.Name != "After" {
		t.Errorf("Incorrect List name '%s'", l.Name)
	}
}

func TestListAfterActionOnArchive(t *testing.T) {
	a := &Action{
		Type: "updateCard",
		Data: &ActionData{
			List:  &List{Name: "SameList"},
			Board: &Board{},
			Card:  &Card{Closed: true},
			Old:   &Card{Closed: false},
		},
	}
	l := ListAfterAction(a)
	if l != nil {
		t.Error("ListAfterAction() should be nil after an archive.")
	}
}

func TestListAfterActionOnUnarchive(t *testing.T) {
	a := &Action{
		Type: "updateCard",
		Data: &ActionData{
			List:  &List{Name: "SameList"},
			Board: &Board{},
			Card:  &Card{Closed: false},
			Old:   &Card{Closed: true},
		},
	}
	l := ListAfterAction(a)
	if l == nil {
		t.Error("ListAfterAction() should not be nil after an unarchive.")
	}
	if l.Name != "SameList" {
		t.Errorf("Incorrect List name '%s'.", l.Name)
	}
}

func TestListAfterActionOnCopyCard(t *testing.T) {
	a := &Action{
		Type: "copyCard",
		Data: &ActionData{
			List:  &List{Name: "FirstList"},
			Board: &Board{},
			Card:  &Card{Closed: false},
		},
	}
	l := ListAfterAction(a)
	if l == nil {
		t.Error("ListAfterAction() should not be nil after a copy")
	}
	if l.Name != "FirstList" {
		t.Errorf("Incorrect List name '%s'.", l.Name)
	}
}
