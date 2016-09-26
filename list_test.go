package trello

import (
	"testing"
)

func TestGetListsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("lists", "board-lists-api-example.json").URL
	lists, err := board.GetLists(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(lists) != 3 {
		t.Errorf("Expected 1 list, got %d", len(lists))
	}
}
