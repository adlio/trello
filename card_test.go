package trello

import (
	"testing"
)

func TestGetCardsOnBoard(t *testing.T) {
	c := NewClient("user", "pass")
	boardResponse := mockResponse("boards", "cI66RoQS.json")
	cardsResponse := mockResponse("cards", "board-cards-api-example.json")

	c.BaseURL = boardResponse.URL
	board, err := c.GetBoard("cI66RoQs", Defaults)
	if err != nil {
		t.Fatal(err)
	}

	c.BaseURL = cardsResponse.URL
	cards, err := board.GetCards(Defaults)
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(cards))
	}
}

// Utility function to get a simple response from Client.GetCard()
//
func testCard(t *testing.T) *Card {
	c := NewClient("user", "pass")
	c.BaseURL = mockResponse("cards", "card-api-example.json").URL
	card, err := c.GetCard("4eea503", Defaults)
	if err != nil {
		t.Fatal(err)
	}
	return card
}
