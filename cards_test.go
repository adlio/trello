package trello

import (
	"testing"
)

func TestGetBoardCards(t *testing.T) {
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
