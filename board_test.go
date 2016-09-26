package trello

import (
	"testing"
)

func testBoard(t *testing.T) *Board {
	c := NewClient("user", "pass")
	boardResponse := mockResponse("boards", "cI66RoQS.json")
	c.BaseURL = boardResponse.URL
	board, err := c.GetBoard("cIRoQS", Defaults)
	if err != nil {
		t.Fatal(err)
	}
	return board
}

func TestGetBoard(t *testing.T) {
	board := testBoard(t)
	if board.Name != "Trello Public API" {
		t.Errorf("Incorrect board name '%s'", board.Name)
	}

	if board.LabelNames.Green != "Participate!" {
		t.Errorf("Expected Green label 'Participate!'. Got '%s'", board.LabelNames.Green)
	}
}

func TestGetBoards(t *testing.T) {
	c := NewClient("user", "pass")

	memberResponse := mockResponse("members", "api-example.json")
	boardsResponse := mockResponse("boards", "member-boards-example.json")

	c.BaseURL = memberResponse.URL
	member, err := c.GetMember("4ee7df1", Defaults)
	if err != nil {
		t.Fatal(err)
	}

	c.BaseURL = boardsResponse.URL
	boards, err := member.GetBoards(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(boards) != 2 {
		t.Errorf("Expected 2 boards. Got %d", len(boards))
	}

	if boards[0].Name != "Example Board" {
		t.Errorf("Name of first board incorrect. Got: '%s'", boards[0].Name)
	}

	if boards[1].Name != "Public Board" {
		t.Errorf("Name of second board incorrect. Got: '%s'", boards[1].Name)
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
