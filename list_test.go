// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
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

func TestGetListWithCards(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("lists", "list-api-example.json").URL
	list, err := c.GetList("4eea4ff", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Cards) == 0 {
		t.Fatal("cannot test cards as non was available in lists mock response")
	}
	if list.Cards[0].client == nil {
		t.Fatal("client not set on cards")
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

func TestGetListsOnBoardWithCards(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("lists", "board-lists-api-example.json").URL
	lists, err := board.GetLists(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	for i := range lists {
		if len(lists[i].Cards) == 0 {
			t.Fatal("cannot test cards as non was available in lists mock response")
		}
		if lists[i].Cards[0].client == nil {
			t.Fatal("client not set on cards")
		}
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

func TestCreateList(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("lists", "create-list-example.json").URL

	board := Board{
		client: c,
		ID:     "5c41027ca9c378795b5a5036",
	}

	listName := "hello"

	list, err := board.CreateList(listName, Arguments{"pos": "35"})
	if err != nil {
		t.Error(err)
	}
	if list.ID != "5ccd793e91682684235c0b13" {
		t.Errorf("Expected list to pick up an ID. Instead got '%s'.", list.ID)
	}
	if list.IDBoard != "5c41027ca9c378795b5a5036" {
		t.Errorf("Expected list to pick up board ID. Instead got '%s'.", list.IDBoard)
	}
	if list.Name != listName {
		t.Errorf("Expected list to pick up name. Instead got '%s'", list.Name)
	}
	if list.Pos != 35 {
		t.Errorf("Expected the returned list to pick up a position. Instead got '%v'.", list.Pos)
	}
	if list.Closed != false {
		t.Errorf("Expected list to pick up Closed. Instead got '%v'", list.Closed)
	}
	if list.client == nil {
		t.Errorf("Expected list to pick up client. Instead got nil")
	}
}

func TestUpdateList(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("lists", "create-list-example.json").URL

	// preparation: create board
	board := Board{
		client: c,
		ID:     "5c41027ca9c378795b5a5036",
	}

	listName := "hello"

	list, err := board.CreateList(listName, Arguments{"pos": "35"})
	if err != nil {
		t.Error(err)
	}
	if list.ID != "5ccd793e91682684235c0b13" {
		t.Errorf("Expected list to pick up an ID. Instead got '%s'.", list.ID)
	}
	if list.IDBoard != "5c41027ca9c378795b5a5036" {
		t.Errorf("Expected list to pick up board ID. Instead got '%s'.", list.IDBoard)
	}
	if list.Name != listName {
		t.Errorf("Expected list to pick up name. Instead got '%s'", list.Name)
	}
	if list.Pos != 35 {
		t.Errorf("Expected the returned list to pick up a position. Instead got '%v'.", list.Pos)
	}

	// update
	listName = "updated-list-name"

	c.BaseURL = mockResponse("lists", "update-list-example.json").URL
	updateArgs := Arguments{"name": listName, "idBoard": "5d31c3d8615ae32928635a28"}

	err = list.Update(updateArgs)

	if err != nil {
		t.Error(err)
	}
	if list.ID != "5ccd793e91682684235c0b13" {
		t.Errorf("Expected list to pick up an ID. Instead got '%s'.", list.ID)
	}
	if list.IDBoard != "5d31c3d8615ae32928635a28" {
		t.Errorf("Expected list to pick up board ID. Instead got '%s'.", list.IDBoard)
	}
	if list.Name != listName {
		t.Errorf("Expected list to pick up name. Instead got '%s'", list.Name)
	}
	if list.Pos != 24576 {
		t.Errorf("Expected the returned list to pick up a position. Instead got '%v'.", list.Pos)
	}
}
