package trello

import (
	"testing"
)

func TestGetCardsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("cards", "board-cards-api-example.json").URL
	cards, err := board.GetCards(Defaults)
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
	cards, err := list.GetCards(Defaults)
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 1 {
		t.Errorf("Expected 1 cards, got %d", len(cards))
	}
}

// Utility function to get a simple response from Client.GetCard()
//
func testCard(t *testing.T) *Card {
	c := testClient()
	c.BaseURL = mockResponse("cards", "card-api-example.json").URL
	card, err := c.GetCard("4eea503", Defaults)
	if err != nil {
		t.Fatal(err)
	}
	return card
}
