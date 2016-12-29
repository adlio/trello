// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"testing"
	"time"
)

func TestCardCreatedAt(t *testing.T) {
	c := Card{}
	c.ID = "4d5ea62fd76aa1136000000c"
	ts := c.CreatedAt()
	if ts.IsZero() {
		t.Error("Time shouldn't be zero.")
	}
	if ts.Unix() != 1298048559 {
		t.Errorf("Incorrect CreatedAt() time: '%s'.", ts.Format(time.RFC3339))
	}
}

func TestGetCardsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockDynamicPathResponse().URL
	cards, err := board.GetCards(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(cards))
	}
}

func TestGetCardsInList(t *testing.T) {
	list := testList(t)
	list.client.BaseURL = mockResponse("cards", "list-cards-api-example.json").URL
	cards, err := list.GetCards(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 1 {
		t.Errorf("Expected 1 cards, got %d", len(cards))
	}
}

func TestBoardContainsCopyOfCard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("actions", "board-actions-copyCard.json").URL
	firstResult, err := board.ContainsCopyOfCard("57f50c552b96e3fffe588aad", Defaults())
	if err != nil {
		t.Error(err)
	}
	if firstResult {
		t.Errorf("Incorrect Copy test: Card 57f50c552b96e3fffe588aad was never copied.")
	}

	secondResult, err := board.ContainsCopyOfCard("57914873fd2de1a10f3cb422", Defaults())
	if err != nil {
		t.Error(err)
	}
	if !secondResult {
		t.Errorf("ContainsCopyOfCard(57f50c552b96e3fffe588aad) should have been true.")
	}
}

func TestCreateCard(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("cards", "card-create.json").URL
	dueDate := time.Now().AddDate(0, 0, 3)

	card := Card{
		Name:   "Test Card Create",
		Desc:   "What its about",
		Due:    &dueDate,
		IDList: "57f03a06b5ff33a63c8be316",
	}

	err := c.CreateCard(&card, Arguments{"pos": "top"})
	if err != nil {
		t.Error(err)
	}

	if card.Pos != 8192 {
		t.Errorf("Expected card to pick up a new Pos value. Instead got %.2f.", card.Pos)
	}

	if card.DateLastActivity == nil {
		t.Error("Expected card to pick up a last activity date. Was nil.")
	}

	if card.ID != "57f5183c691585658d408681" {
		t.Errorf("Expected card to pick up an ID. Instead got '%s'.", card.ID)
	}
}

func TestAddCardToList(t *testing.T) {
	l := testList(t)
	l.client.BaseURL = mockResponse("cards", "card-posted-to-bottom-of-list.json").URL
	dueDate := time.Now().AddDate(0, 0, 1)

	card := Card{
		Name: "Test Card POSTed to List",
		Desc: "This is its description.",
		Due:  &dueDate,
	}

	err := l.AddCard(&card, Arguments{"pos": "bottom"})
	if err != nil {
		t.Error(err)
	}

	if card.Pos != 32768 {
		t.Errorf("Expected card to pick up a new Pos value. Instead got %.2f.", card.Pos)
	}

	if card.DateLastActivity == nil {
		t.Error("Expected card to pick up a last activity date. Was nil.")
	}

	if card.ID != "57f5118667db8839dab68698" {
		t.Errorf("Expected card to pick up an ID. Instead got '%s'.", card.ID)
	}
}

func TestCopyCardToList(t *testing.T) {
	c := testCard(t)
	c.client.BaseURL = mockResponse("cards", "card-copied.json").URL

	newCard, err := c.CopyToList("57f03a022cd45c863ca581f1", Defaults())
	if err != nil {
		t.Error(err)
	}

	if newCard.ID == c.ID {
		t.Errorf("New card should have a new ID: '%s'.", newCard.ID)
	}

	if newCard.Pos != 16384 {
		t.Errorf("Expected new card to have correct Pos value. Got %.2f", newCard.Pos)
	}
}

func TestGetParentCard(t *testing.T) {
	c := testCard(t)
	c.client.BaseURL = mockDynamicPathResponse().URL

	parent, err := c.GetParentCard(Defaults())
	if err != nil {
		t.Error(err)
	}
	if parent == nil {
		t.Errorf("Problem")
	}
}

func TestGetAncestorCards(t *testing.T) {
	c := testCard(t)
	c.client.BaseURL = mockDynamicPathResponse().URL

	ancestors, err := c.GetAncestorCards(Defaults())
	if err != nil {
		t.Error(err)
	}
	if len(ancestors) != 1 {
		t.Errorf("Expected 1 ancestor, got %d", len(ancestors))
	}
}

// Utility function to get a simple response from Client.GetCard()
//
func testCard(t *testing.T) *Card {
	c := testClient()
	c.BaseURL = mockResponse("cards", "card-api-example.json").URL
	card, err := c.GetCard("4eea503", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return card
}
