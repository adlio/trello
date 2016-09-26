package trello

import (
	"testing"
)

func TestGetActionsOnBoard(t *testing.T) {
	c := NewClient("user", "pass")

	boardResponse := mockResponse("boards", "cI66RoQS.json")
	actionsResponse := mockResponse("actions", "board-actions-api-example.json")

	c.BaseURL = boardResponse.URL
	board, err := c.GetBoard("cIRoQS", Defaults)
	if err != nil {
		t.Fatal(err)
	}

	c.BaseURL = actionsResponse.URL
	actions, err := board.GetActions(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(actions) != 4 {
		t.Errorf("Expected 1 actions, got %d", len(actions))
	}
}
