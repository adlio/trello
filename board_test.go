package trello

import (
	"testing"
)

func TestGetBoard(t *testing.T) {
	m := mockResponse("boards", "cI66RoQS.json")
	c := NewClient("user", "pass")
	c.BaseURL = m.URL

	board, err := c.GetBoard("cI66RoQS", Defaults)
	if err != nil {
		t.Fatal(err)
	}
	if board == nil {
		t.Error("Board retrieved from c.GetBoard() shouldn't be nil")
	}

	if board.Name != "Trello Public API" {
		t.Errorf("Incorrect board name '%s'", board.Name)
	}

	if board.LabelNames.Green != "Participate!" {
		t.Errorf("Expected Green label 'Participate!'. Got '%s'", board.LabelNames.Green)
	}

}

func TestGetUnauthorizedBoard(t *testing.T) {
	m := mockErrorResponse(401)
	c := NewClient("user", "pass")
	c.BaseURL = m.URL

	_, err := c.GetBoard("boardid", Defaults)
	if err == nil {
		t.Error("GetBoard() should have failed")
	}
}
