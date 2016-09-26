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
		t.Errorf("Expected 1 actions, got %d", len(actions))
	}
}
